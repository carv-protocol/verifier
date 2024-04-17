package data

import (
	"context"

	"github.com/carv-protocol/verifier/internal/biz"
	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/pkg/dblogger"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gozelus/gormotel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewVerifierRepo,
	NewTransactionRepo,
	NewReportTeeAttestationEventRepo,
	wire.Bind(new(biz.VerifierRepo), new(*VerifierRepo)),
	wire.Bind(new(biz.TransactionRepo), new(*TransactionRepo)),
	wire.Bind(new(biz.ReportTeeAttestationEventRepo), new(*ReportTeeAttestationEventRepo)),
)

type TransactionManager interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type contextTxKey struct{}

// Data .
type Data struct {
	cf       *conf.Bootstrap
	db       *gorm.DB
	cleanups []func()
}

func NewData(bootstrap *conf.Bootstrap, logger *log.Helper) (*Data, func()) {
	d := &Data{
		cf: bootstrap,
	}

	if bootstrap.Data.Database != nil {
		db, err := gorm.Open(mysql.Open(bootstrap.Data.Database.Source), &gorm.Config{
			Logger:         dblogger.NewTracer(logger),
			TranslateError: true,
		})
		if err != nil {
			panic(err)
		}
		// report db tracing to open-telemetry
		if err = db.Use(gormotel.Plugin); err != nil {
			panic(err)
		}
		d.db = db
		sqlDB, err := d.db.DB()
		if err != nil {
			panic(err)
		}

		// configure pool parameters
		sqlDB.SetMaxOpenConns(int(bootstrap.Data.Database.MaxOpenConnections))
		sqlDB.SetMaxIdleConns(int(bootstrap.Data.Database.MaxIdleConnections))
		sqlDB.SetConnMaxLifetime(bootstrap.Data.Database.ConnMaxLifetime.AsDuration())
		sqlDB.SetConnMaxIdleTime(bootstrap.Data.Database.ConnMaxIdleTime.AsDuration())

		d.AppendCleanup(func() {
			_ = sqlDB.Close()
			log.Info("db closed")
		})

		// Auto migrate tables
		err = d.db.AutoMigrate(
			&biz.Transaction{},
			&biz.ReportTeeAttestationEvent{},
		)
		if err != nil {
			log.Info("auto migrate tables failed")
			return nil, nil
		}
	}

	return d, func() {
		for _, cleanup := range d.cleanups {
			cleanup()
		}
	}
}
func (d *Data) AppendCleanup(cleanup func()) {
	d.cleanups = append(d.cleanups, cleanup)
}

// ExecTx transaction manager's core func, implemented it by gorm and context value func now
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB try to get transaction handle by the context,
// if the handle do not exist, will return data's DB instance directly
// so do not forget do some sql without withContext func
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}
