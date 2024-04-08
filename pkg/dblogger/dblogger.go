package dblogger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	glog "log"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _ logger.Interface = (*Tracer)(nil)

type Tracer struct {
	slowSqlTime time.Duration
	logger      *log.Helper
}

func NewTracer(logger *log.Helper) *Tracer {
	return &Tracer{logger: logger}
}

func (d *Tracer) LogMode(level logger.LogLevel) logger.Interface {
	return d
}

func (d *Tracer) Info(ctx context.Context, s string, i ...interface{}) {
	d.logger.WithContext(ctx).Debugf(s, i)
}

func (d *Tracer) Warn(ctx context.Context, s string, i ...interface{}) {
	d.logger.WithContext(ctx).Warnf(s, i)
}

func (d *Tracer) Error(ctx context.Context, s string, i ...interface{}) {
	d.logger.WithContext(ctx).Errorf(s, i)
}

func (d *Tracer) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil:
		sql, rows := fc()
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			d.logger.WithContext(ctx).Errorf("[%.3fms] [rows:%v] %s err for %s", float64(elapsed.Nanoseconds())/1e6, rows, sql, err)
		}
	case elapsed > d.slowSqlTime && d.slowSqlTime != 0:
		sql, rows := fc()
		d.logger.WithContext(ctx).Warnf("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	default:
		sql, rows := fc()
		d.logger.WithContext(ctx).Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
}

var _ log.Logger = (*StdLogger)(nil)

type StdLogger struct {
	log  *glog.Logger
	pool *sync.Pool
}

// NewStdLogger new a logger with writer.
func NewStdLogger(w io.Writer) log.Logger {
	return &StdLogger{
		log: glog.New(w, "", 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

// Log print the kv pairs log.
func (l *StdLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}
	buf := l.pool.Get().(*bytes.Buffer)

	keyvals = append(keyvals, "level", level.String())
	// json 化
	v := make(map[string]interface{})
	for i := 0; i < len(keyvals); i += 2 {
		v[keyvals[i].(string)] = keyvals[i+1]
	}
	contentBs, _ := json.Marshal(v)
	fmt.Fprintf(buf, string(contentBs))
	l.log.Output(4, buf.String())
	buf.Reset()
	l.pool.Put(buf)
	return nil
}
