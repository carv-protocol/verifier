package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"

func main() {

	a := app.New()
	myWindow := a.NewWindow("VerifierClient")
	myCanvas := myWindow.Canvas()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
			dialog.ShowInformation("Help", "This is a simple verifier client", myWindow)
		}),
	)

	Title := widget.NewLabel("Welcome To Use Verifier Client")
	TitleCss(Title)
	// form
	RpcUrl := widget.NewEntry()
	RpcUrl.SetPlaceHolder("https://opbnb-testnet-rpc.bnbchain.org")
	PrivateKey := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Entry", Widget: RpcUrl},
			{Text: "test", Widget: PrivateKey},
		},
		OnSubmit: func() {
			if RpcUrl.Text == "" {
				//runVerifier(RpcUrl.Text, PrivateKey.Text)
				go runVerifier(RpcUrl.PlaceHolder, PrivateKey.Text)
			} else {
				go runVerifier(RpcUrl.Text, PrivateKey.Text)
			}

			time.Sleep(5 * time.Second)
			data := checkVerifierIsActive()
			if data.Code == 200 {
				dialog.ShowInformation("Verifier Running", data.Msg, myWindow)
			}
		},
	}

	status := widget.NewLabel("Status")
	go showStatus(status)

	myCanvas.SetContent(container.NewVBox(toolbar, Title, form, status))
	myWindow.Resize(fyne.NewSize(600, 400))

	myWindow.ShowAndRun()
	myWindow.Close()
}

func TitleCss(title *widget.Label) {
	title.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	title.Alignment = fyne.TextAlignCenter

}

func showStatus(statusLabel *widget.Label) {
	statusLabel.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	for {
		time.Sleep(2 * time.Second)
		data := checkVerifierIsActive()
		fmt.Println(data.Msg)
		statusLabel.SetText(data.Msg)
	}
}
