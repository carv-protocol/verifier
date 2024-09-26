package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	originHttp "net/http"
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	commonInternal "github.com/carv-protocol/verifier/internal/common"
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
	CommissionRate float64
}

func init() {
	flag.StringVar(&flagVar.Conf, "conf", "", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagVar.PrivateKey, "private-key", "", "private key, eg: -private-key 9a8bd8c....21dec")
	flag.StringVar(&flagVar.KeystorePath, "keystore-path", "", "keystore path, eg: -keystore-path . default: .")
	flag.StringVar(&flagVar.KeystorePassword, "keystore-password", "", "keystore password, eg: -keystore-password 123456")
	flag.BoolVar(&flagVar.GenerateKeystore, "generate-keystore", false, "generate keystore, eg: -generate-keystore")
	flag.StringVar(&flagVar.RewardAddress, "reward-address", "", "reward address, eg: -reward-address 0x123456")
	flag.Float64Var(&flagVar.CommissionRate, "commission-rate", 0, "commission rate, eg: -commission-rate 10")
}

func newApp(logger log.Logger, workerServer *worker.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(workerServer),
	)
}

func main() {
	flag.Parse()
	if flagVar.GenerateKeystore {
		account := key_manager.GenerateKeystore(flagVar.KeystorePath)
		fmt.Printf("generate keystore success, Please delegate to: %v, and run again.", account.Address.Hex())
		return
	}

	// By default use local path, easier for executable file
	configFile := ""
	if flagVar.Conf != "" {
		configFile = flagVar.Conf
	} else {
		configFile = fetchConfigFromURL(commonInternal.BASE_URl + "config_local.yaml")
	}
	c := config.New(
		config.WithSource(
			file.NewSource(configFile),
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

	if err := key_manager.NewKeyManager(bc.Wallet, flagVar.PrivateKey, flagVar.KeystorePath, flagVar.KeystorePassword); err != nil {
		panic(err)
	}

	if flagVar.Conf != "" {
		flagVar.RewardAddress = bc.Wallet.RewardClaimerAddr
		flagVar.CommissionRate = bc.Wallet.CommissionRate
	}
	if common.IsHexAddress(flagVar.RewardAddress) == false {
		panic("invalid reward address")
	}

	if flagVar.CommissionRate < 0 || flagVar.CommissionRate > 100 {
		panic("invalid commission rate")
	}
	bc.Wallet.RewardClaimerAddr = flagVar.RewardAddress
	bc.Wallet.CommissionRate = flagVar.CommissionRate * 100 // 1% -> 100/10000

	// Commission should be a whole number if input with at most 2 deciamls
	// Should be in range 0 - 10000. If input value is 1.5, the new commission will be 150.
	if bc.Wallet.CommissionRate != math.Trunc(bc.Wallet.CommissionRate) {
		panic("invalid commission rate, at most 2 decimals")
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

func fetchConfigFromURL(url string) string {
	resp, err := originHttp.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)

	// save as temp
	tmpFile, err := os.CreateTemp("", "config_local*.yaml")
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	if _, err := tmpFile.Write(data); err != nil {
		panic(err)
	}
	return tmpFile.Name()

}
