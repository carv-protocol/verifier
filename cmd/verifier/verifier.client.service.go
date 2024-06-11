package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	fyneV2 "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"github.com/carv-protocol/verifier/internal/common"
	"github.com/carv-protocol/verifier/internal/conf"
	"github.com/carv-protocol/verifier/internal/key_manager"
	"github.com/carv-protocol/verifier/internal/worker"
	"github.com/carv-protocol/verifier/pkg/stdlogger"
	"github.com/djherbis/times"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fsnotify/fsnotify"
	kratosV2 "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/shirou/gopsutil/cpu"
	"io"
	originHttp "net/http"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	RpcUrl string
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
)

func init() {
	flag.StringVar(&flagVar.Conf, "conf", "configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagVar.PrivateKey, "private-key", "", "private key, eg: -private-key 9a8bd8c....21dec")
	flag.StringVar(&flagVar.KeystorePath, "keystore-path", "", "keystore path, eg: -keystore-path .")
	flag.StringVar(&flagVar.KeystorePassword, "keystore-password", "", "keystore password, eg: -keystore-password 123456")
	flag.BoolVar(&flagVar.GenerateKeystore, "generate-keystore", false, "generate keystore, eg: -generate-keystore")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, workerServer *worker.Server) *kratosV2.App {
	return kratosV2.New(
		kratosV2.ID(id),
		kratosV2.Name(Name),
		kratosV2.Version(Version),
		kratosV2.Metadata(map[string]string{}),
		kratosV2.Logger(logger),
		kratosV2.Server(
			gs,
			hs,
			workerServer,
		),
	)
}

type Response struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

func runVerifier(rpcUrl, privateKey string, w fyneV2.Window) error {
	// store rpcUrl
	RpcUrl = rpcUrl

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
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		data := &Response{
			Code:   200,
			Result: "",
			Msg:    "No logs found",
		}
		return data

	}
	files, err := os.ReadDir("logs")
	if err != nil {
		panic(err)
	}
	sort.Slice(files, func(i, j int) bool {
		timeI, errI := times.Stat("logs/" + files[i].Name())
		timeJ, errJ := times.Stat("logs/" + files[j].Name())
		if errI != nil || errJ != nil {
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

// get System information
func getSystemInformation(sys *canvas.Text, allocText *canvas.Text, totalAlloc *canvas.Text, cpuUsage *canvas.Text) {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			sys.Text = "System: " + runtime.GOOS + " " + runtime.GOARCH + " " + strconv.Itoa(runtime.NumCPU()) + " CPU(s)"
			sys.Refresh()

			allocMB := float64(memStats.Alloc) / 1024 / 1024
			allocText.Text = "Allocated memory: " + strconv.FormatFloat(allocMB, 'f', 2, 64) + " MB"
			allocText.Refresh()

			totalAllocMB := float64(memStats.TotalAlloc) / 1024 / 1024
			totalAlloc.Text = "Total Allocated memory: " + strconv.FormatFloat(totalAllocMB, 'f', 2, 64) + " MB"
			totalAlloc.Refresh()

			cpuPersent, _ := cpu.Percent(time.Second, false)
			cpuUsage.Text = "CPU Usage: " + strconv.FormatFloat(cpuPersent[0], 'f', 2, 64) + "%"
			cpuUsage.Refresh()
		}

	}()
}
func fetchConfigFromURL(url string) (string, error) {
	resp, err := originHttp.Get(url)
	if err != nil {
		return "", err
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
		return "", err
	}
	defer tmpFile.Close()
	if _, err := tmpFile.Write(data); err != nil {
		return "", err
	}
	return tmpFile.Name(), nil

}

// get Latest Block information from log file

func logWatcher(label *canvas.Text) error {

	_, err := os.ReadDir("logs")
	if err != nil {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			return err
		}
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					//fmt.Println("modified logFile:", event.Name)
					blockHeight := getLatestBlockHeight(event.Name)
					if blockHeight != 0 {
						label.Text = "Current Block Height: " + strconv.Itoa(blockHeight)
						label.Refresh()
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	logFile, err := getLatestLogFile()
	if err != nil {
		return err
	}
	err = watcher.Add("logs/" + logFile.Name())
	if err != nil {
		return err
	}
	<-done
	return nil
}

type LogEntry struct {
	Msg string `json:"msg"`
}

func getLatestBlockHeight(filename string) int {
	logFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR", err)
		return 0
	}
	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	var lastLine string
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	var logEntry LogEntry
	json.Unmarshal([]byte(lastLine), &logEntry)

	// Assuming the block height is in the format "start block 30524656"
	parts := strings.Split(logEntry.Msg, " ")
	blockHeight := 0
	if len(parts) > 5 {
		blockHeight, err = strconv.Atoi(parts[5][:len(parts[5])-1])
		if err != nil {
			fmt.Println("ERROR", err)
		}
	}

	return blockHeight
}

// utils
func getLatestLogFile() (os.DirEntry, error) {
	files, err := os.ReadDir("logs")
	if err != nil {
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool {
		timeI, errI := times.Stat("logs/" + files[i].Name())
		timeJ, errJ := times.Stat("logs/" + files[j].Name())
		if errI != nil || errJ != nil {
			return false
		}
		return timeI.ModTime().After(timeJ.ModTime())
	})
	if len(files) == 0 {
		return nil, fmt.Errorf("No logs found")
	}

	return files[0], nil
}

// get latest block from chain
func getLatestBlockFromChain(label *canvas.Text) {
	for {
		time.Sleep(1 * time.Second)
		if RpcUrl == "" {
			RpcUrl = common.OP_BNB_RPC_URL
		}
		client, err := ethclient.Dial(RpcUrl)
		if err != nil {
			log.Fatal(err)
		}

		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		label.Text = "Latest Block Height: " + header.Number.String()
		label.Refresh()
	}

	//fmt.Println(header.Number.String()) // prints latest block number
}