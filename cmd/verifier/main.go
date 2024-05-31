package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
func main() {

	a := app.NewWithID("com.verifier.client")
	myWindow := a.NewWindow("VerifierClient")
	myWindow.SetMainMenu(makeMenu(a, myWindow))
	myWindow.SetMaster()

	Title := widget.NewLabel("Welcome To Use Verifier Client")
	TitleCss(Title)

	// verifier status
	status := newColoredLabel("InActive", color.RGBA{255, 0, 0, 255})
	go showStatus(status)

	//
	content := container.NewBorder(container.NewHBox(layout.NewSpacer(), Title, layout.NewSpacer()), container.NewHBox(layout.NewSpacer(), status), nil, nil)

	content.Resize(fyne.NewSize(800, 600))
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.ShowAndRun()
	myWindow.Close()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	fileMenu := fyne.NewMenu("Run",
		fyne.NewMenuItem("start", func() {
			newWindow := a.NewWindow("RunNode")
			// form
			RpcUrl := widget.NewEntry()
			RpcUrl.SetPlaceHolder("https://opbnb-testnet-rpc.bnbchain.org")
			PrivateKey := widget.NewPasswordEntry()
			form := &widget.Form{
				Items: []*widget.FormItem{
					{Text: "RpcUrl", Widget: RpcUrl},
					{Text: "PrivateKey", Widget: PrivateKey},
				},
				OnSubmit: func() {
					if RpcUrl.Text == "" {
						//runVerifier(RpcUrl.Text, PrivateKey.Text)
						go runVerifier(RpcUrl.PlaceHolder, PrivateKey.Text, newWindow)
					} else {
						go runVerifier(RpcUrl.Text, PrivateKey.Text, newWindow)
					}

					time.Sleep(5 * time.Second)
					data := checkVerifierIsActive()
					if data.Code == 200 {
						dialog.ShowInformation("Verifier Running", data.Msg, newWindow)
					} else {
						dialog.ShowError(fmt.Errorf("Verifier not Running"), newWindow)
					}
				},
			}
			newWindow.Resize(fyne.NewSize(600, 400))
			newWindow.SetContent(form)
			newWindow.Show()
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Quit", func() { dialog.ShowInformation("Coming", "Coming", w) }),
	)

	editMenu := fyne.NewMenu("Log",
		fyne.NewMenuItem("Check", func() {
			newWindow := a.NewWindow("LogWindow")
			logsTextArea := widget.NewMultiLineEntry()
			setTextArea(logsTextArea)
			go showLogs(logsTextArea)
			newWindow.Resize(fyne.NewSize(600, 400))
			logsTextArea.Resize(fyne.NewSize(600, 400))
			newWindow.SetContent(logsTextArea)
			newWindow.Show()
		}),
	)
	//windowMenu := fyne.NewMenu("Windows")
	menus := make([]*fyne.Menu, 0)
	menus = append(menus, fileMenu, editMenu)
	mainMenu := &fyne.MainMenu{
		Items: menus,
	}
	return mainMenu
}

func TitleCss(title *widget.Label) {
	title.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	title.Alignment = fyne.TextAlignCenter

}

func setTextArea(textArea *widget.Entry) {
	textArea.Resize(fyne.NewSize(1000, 800))
}

func showStatus(statusLabel *coloredLabel) {
	for {
		time.Sleep(2 * time.Second)
		data := checkVerifierIsActive()
		fmt.Println(data.Msg)
		if data.Code == 200 {
			statusLabel.Text = "Active"
			statusLabel.Color = color.RGBA{0, 255, 0, 255}
		} else {
			statusLabel.Color = color.RGBA{255, 0, 0, 255}
		}

		statusLabel.Refresh()
	}
}

func showLogs(textArea *widget.Entry) {
	textArea.SetText("")
	logsData := getSystemLogs()
	textArea.SetText(logsData.Result)

}

// canvas text extension
type coloredLabel struct {
	widget.BaseWidget
	Text  string
	Color color.Color
}

func newColoredLabel(text string, color color.Color) *coloredLabel {
	l := &coloredLabel{
		Text:  text,
		Color: color,
	}
	l.ExtendBaseWidget(l)
	return l
}

func (l *coloredLabel) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(l.Text, l.Color)
	text.Alignment = fyne.TextAlignLeading

	return &coloredLabelRenderer{
		label: l,
		text:  text,
	}
}

type coloredLabelRenderer struct {
	label *coloredLabel
	text  *canvas.Text
}

func (r *coloredLabelRenderer) Layout(size fyne.Size) {
	r.text.Resize(size)
}

func (r *coloredLabelRenderer) MinSize() fyne.Size {
	return r.text.MinSize()
}

func (r *coloredLabelRenderer) Refresh() {
	r.text.Text = r.label.Text
	r.text.Color = r.label.Color
	r.text.Refresh()
}

func (r *coloredLabelRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *coloredLabelRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.text}
}

func (r *coloredLabelRenderer) Destroy() {}
