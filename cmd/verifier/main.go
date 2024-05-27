package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/carv-protocol/verifier/internal/common"
	"github.com/carv-protocol/verifier/pkg/stdlogger"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
	"io"
	originHttp "net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/key_manager"
	"github.com/carv-protocol/verifier/internal/worker"
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
	//go:embed assets/*
	f embed.FS
)

type FlagVar struct {
	Conf             string
	PrivateKey       string
	KeystorePath     string
	KeystorePassword string
	GenerateKeystore bool
}

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

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exeDir := filepath.Dir(exePath)
	fmt.Println(exeDir)

	router := mux.NewRouter()

	router.PathPrefix("/assets/").Handler(originHttp.FileServer(originHttp.FS(f)))
	router.HandleFunc("/verifier/run", runVerifier).Methods("POST")

	httpSrv := http.NewServer(http.Address(":8080"))
	httpSrv.HandlePrefix("/", router)

	app2 := kratos.New(
		kratos.Name("static"),
		kratos.Server(httpSrv),
	)
	go openBrower("http://localhost:8080/assets/")

	fmt.Println(":8080")
	if err = app2.Run(); err != nil {
		log.Fatal(err)
	}

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

func openBrower(url string) {

	time.Sleep(100 * time.Millisecond)
	if err := browser.OpenURL(url); err != nil {
		fmt.Println("open browser error", err)
	}
}

type Setting struct {
	RpcUrl     string   `json:"rpcUrl"`
	PrivateKey string   `json:"privateKey"`
	Urls       []string `json:"urls,omitempty"`
}

func runVerifier(w originHttp.ResponseWriter, r *originHttp.Request) {
	var setting Setting

	err := json.NewDecoder(r.Body).Decode(&setting)
	if err != nil {
		originHttp.Error(w, err.Error(), originHttp.StatusBadRequest)
		return
	}
	fmt.Printf("setting: %+v\n", setting)

	tempFile, tempFileErr := fetchConfigFromURL(common.BASE_URl + "config_local.yaml")
	if tempFileErr != nil {
		panic(tempFileErr)
	}
	c := config.New(
		config.WithSource(
			file.NewSource(tempFile),
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

	// update rpc and privateKey
	bc.Chain.RpcUrl = setting.RpcUrl
	bc.Wallet.Mode = 1
	bc.Wallet.PrivateKey = setting.PrivateKey

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
		w.WriteHeader(originHttp.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		data := map[string]interface{}{
			"message": "Verifier Run Failed",
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(jsonData)
		if err != nil {
			return
		}
		return
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
