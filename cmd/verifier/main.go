package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/browser"
	originHttp "net/http"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/carv-protocol/verifier/internal/key_manager"
)

// go build -ldflags "-X main.Version=x.y.z"

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

	// start exec file
	//err := os.Chmod("./verifierclient-Setup-1.0.0.exe", 0777)
	//if err != nil {
	//	log.Fatal("chmod : ", err)
	//}
	//cmd := exec.Command("./verifierclient-Setup-1.0.0.exe")
	//err = cmd.Start()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//go func() {
	//	err = cmd.Wait()
	//	if err != nil {
	//		log.Fatal("Command finished with error: %v", err)
	//	}
	//}()

	router := mux.NewRouter()

	router.PathPrefix("/assets/").Handler(originHttp.FileServer(originHttp.FS(f)))
	router.HandleFunc("/verifier/run", runVerifier).Methods("POST")
	router.HandleFunc("/verifier/logs", getSystemLogs).Methods("GET")

	httpSrv := http.NewServer(http.Address(":8080"))
	httpSrv.HandlePrefix("/", router)

	app2 := kratos.New(
		kratos.Name("static"),
		kratos.Server(httpSrv),
	)
	//go openBrower("http://localhost:8080/assets/")

	if err := app2.Run(); err != nil {
		log.Fatal(err)
	}

}

func openBrowner(url string) {
	time.Sleep(100 * time.Millisecond)
	if err := browser.OpenURL(url); err != nil {
		fmt.Println("open browser error", err)
	}
}
