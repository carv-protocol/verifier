package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/common"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/key_manager"
	"github.com/carv-protocol/verifier/internal/worker"
	"github.com/carv-protocol/verifier/pkg/stdlogger"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagVar
	flagVar FlagVar

	id, _ = os.Hostname()
)

type FlagVar struct {
	Conf             string
	PrivateKey       string
	KeystorePath     string
	KeystorePassword string
	GenerateKeystore bool

	RewardAddress  string
	CommissionRate int
}

func init() {
	flag.StringVar(&flagVar.Conf, "conf", "configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagVar.PrivateKey, "private-key", "", "private key, eg: -private-key 9a8bd8c....21dec")
	flag.StringVar(&flagVar.KeystorePath, "keystore-path", "", "keystore path, eg: -keystore-path .")
	flag.StringVar(&flagVar.KeystorePassword, "keystore-password", "", "keystore password, eg: -keystore-password 123456")
	flag.BoolVar(&flagVar.GenerateKeystore, "generate-keystore", false, "generate keystore, eg: -generate-keystore")
	flag.StringVar(&flagVar.RewardAddress, "reward-address", "", "reward address, eg: -reward-address 0x123456")
	flag.IntVar(&flagVar.CommissionRate, "commission-rate", 0, "commission rate, eg: -commission-rate 10")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, workerServer *worker.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
			workerServer,
		),
	)
}

func main() {
	flag.Parse()

	if flagVar.GenerateKeystore {
		key_manager.GenerateKeystore(flagVar.KeystorePath)
		return
	}

	// By default use local path, easier for executable file
	if flagVar.Conf == "" {
		flagVar.Conf = "config_local.yaml"
	}

	c := config.New(
		config.WithSource(
			file.NewSource(flagVar.Conf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	logLevels := []log.Level{
		log.LevelInfo,
		log.LevelWarn,
		log.LevelError,
		log.LevelFatal,
	}
	if bc.Env == conf.ENV_ENV_LOCAL || bc.Env == conf.ENV_ENV_DEV {
		logLevels = append(logLevels, log.LevelDebug)
	}

	logFormat := log.With(stdlogger.NewStdLogger(os.Stdout, logLevels...),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	logger := log.NewHelper(logFormat)
	//TODO Check Reward Address and Commission rate

	if err := key_manager.NewKeyManager(bc.Wallet, flagVar.PrivateKey, flagVar.KeystorePath, flagVar.KeystorePassword); err != nil {
		panic(err)
	}

	if common.IsHexAddress(flagVar.RewardAddress) == false {
		panic("invalid reward address")
	}
	if flagVar.CommissionRate < 0 || flagVar.CommissionRate > 100 {
		panic("invalid commission rate")
	}
	app, cleanup, err := wireApp(&bc, logFormat, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
