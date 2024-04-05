package stdlogger

import (
	"bytes"
	"encoding/json"
	"fmt"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
	"io"
	"log"
	"sync"
)

var _ klog.Logger = (*stdLogger)(nil)
var allLevel = []klog.Level{klog.LevelDebug, klog.LevelInfo, klog.LevelWarn, klog.LevelError, klog.LevelFatal}

type stdLogger struct {
	log    *log.Logger
	pool   *sync.Pool
	levels []klog.Level
}

// NewStdLogger new a logger with writer.
func NewStdLogger(w io.Writer, levels ...klog.Level) klog.Logger {
	if len(levels) == 0 {
		levels = allLevel
	}
	return &stdLogger{
		levels: levels,
		log:    log.New(w, "", 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

// Log print the kv pairs klog.
func (l *stdLogger) Log(level klog.Level, keyvals ...interface{}) error {
	if !lo.Contains(l.levels, level) {
		return nil
	}
	if len(keyvals) == 0 {
		return nil
	}
	if (len(keyvals) & 1) == 1 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}
	buf := l.pool.Get().(*bytes.Buffer)

	// Start with the log level
	data := map[string]interface{}{"level": level.String()}

	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			key = fmt.Sprintf("%v", keyvals[i])
		}
		data[key] = keyvals[i+1]
	}

	encodedData, err := json.Marshal(data)
	if err != nil {
		// Handle error: you can decide whether you want to return the error or handle it differently.
		return err
	}

	buf.Write(encodedData)
	_ = l.log.Output(4, buf.String())
	buf.Reset()
	l.pool.Put(buf)
	return nil
}

func (l *stdLogger) Close() error {
	return nil
}
