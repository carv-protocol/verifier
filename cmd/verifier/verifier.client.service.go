package main

import (
	"embed"
	"flag"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/carv-protocol/verifier/internal/common"
	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/key_manager"
	"github.com/carv-protocol/verifier/internal/worker"
	"github.com/carv-protocol/verifier/pkg/stdlogger"
	"github.com/djherbis/times"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	originHttp "net/http"
	"os"
	"sort"
)

type FlagVar struct {
	Conf             string
	PrivateKey       string
	KeystorePath     string
	KeystorePassword string
	GenerateKeystore bool
}

var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagVar
	flagVar FlagVar

	id, _ = os.Hostname()
	//go:embed assets/*
	f embed.FS
)

func init() {
	flag.StringVar(&flagVar.Conf, "conf", "configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagVar.PrivateKey, "private-key", "", "private key, eg: -private-key 9a8bd8c....21dec")
	flag.StringVar(&flagVar.KeystorePath, "keystore-path", "", "keystore path, eg: -keystore-path .")
	flag.StringVar(&flagVar.KeystorePassword, "keystore-password", "", "keystore password, eg: -keystore-password 123456")
	flag.BoolVar(&flagVar.GenerateKeystore, "generate-keystore", false, "generate keystore, eg: -generate-keystore")
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

type VerifierClientService interface {
}

type Setting struct {
	RpcUrl     string `json:"rpcUrl"`
	PrivateKey string `json:"privateKey"`
}

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

func runVerifier(rpcUrl, privateKey string, w fyne.Window) error {
	flag.Parse()

	if flagVar.GenerateKeystore {
		key_manager.GenerateKeystore(flagVar.KeystorePath)
	}

	tempFile, tempFileErr := fetchConfigFromURL(common.BASE_URl + "config_local.yaml")
	if tempFileErr != nil {
		dialog.ShowError(tempFileErr, w)
		return tempFileErr
	}
	c := config.New(
		config.WithSource(
			file.NewSource(tempFile),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		dialog.ShowError(err, w)
		return err
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		dialog.ShowError(err, w)
		return err
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

	// update rpc and privateKey
	bc.Chain.RpcUrl = rpcUrl
	bc.Wallet.Mode = 1
	bc.Wallet.PrivateKey = privateKey

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
		dialog.ShowError(err, w)
		return err
	}

	app, cleanup, err := wireApp(&bc, logFormat, logger)
	if err != nil {
		dialog.ShowError(err, w)
		return err
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		dialog.ShowError(err, w)
		return err
	}
	return nil
}

// Get System log from logs file

func getSystemLogs() *Response {
	// get log direction file
	files, err := os.ReadDir("logs")
	if err != nil {
		panic(err)
	}
	sort.Slice(files, func(i, j int) bool {
		timeI, errI := times.Stat("logs/" + files[i].Name())
		timeJ, errJ := times.Stat("logs/" + files[j].Name())
		if errI != nil || errJ != nil {
			fmt.Println(errJ, errI)
			return false
		}
		return timeI.ModTime().After(timeJ.ModTime())
	})
	if len(files) == 0 {
		data := &Response{
			Code:   200,
			Result: "",
			Msg:    "No logs found",
		}
		return data
	}
	logFile, err := os.Open("logs/" + files[0].Name())
	if err != nil {
		panic(err)

	}
	defer logFile.Close()
	data, err := io.ReadAll(logFile)
	if err != nil {
		panic(err)

	}

	respData := &Response{
		Code:   200,
		Result: string(data),
		Msg:    "Logs found",
	}

	return respData
}

func checkVerifierIsActive() *Response {
	verifierRes, err := originHttp.Get("http://127.0.0.1:8000/verifier/helloworld")
	if err != nil {
		data := &Response{
			Code:   500,
			Result: "",
			Msg:    "Verifier is not Active",
		}
		return data
	}
	defer verifierRes.Body.Close()

	respData := &Response{
		Code:   200,
		Result: "",
		Msg:    "Verifier is Active",
	}
	return respData

}
func fetchConfigFromURL(url string) (string, error) {
	resp, err := originHttp.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	// save as temp
	tmpFile, err := os.CreateTemp("", "config_local*.yaml")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()
	if _, err := tmpFile.Write(data); err != nil {
		return "", err
	}
	return tmpFile.Name(), nil

}
