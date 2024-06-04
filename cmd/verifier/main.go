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
	"github.com/carv-protocol/verifier/data"
	"image/color"
	"time"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
func main() {
	// information channel
	stopGetLogsChan := make(chan bool)
	a := app.NewWithID("com.verifier.client")
	myWindow := a.NewWindow("VerifierClient")
	myWindow.SetMainMenu(makeMenu(a, myWindow, stopGetLogsChan))
	myWindow.SetMaster()

	Title := widget.NewLabel("Welcome To Use Verifier Client")
	TitleCss(Title)

	// verifier status
	status := newColoredLabel("InActive", color.RGBA{255, 0, 0, 255})
	go showStatus(status)

	// left
	leftRectangle := canvas.NewRectangle(nil)
	leftRectangle.SetMinSize(fyne.NewSize(400, 600))
	sysText := newColoredLabel("System", color.RGBA{0, 200, 0, 255})
	allocMemoryText := newColoredLabel("Allocated Memory: 0.0", color.RGBA{255, 0, 0, 255})
	totalMemoryText := newColoredLabel("Total Memory: 0.0", color.RGBA{0, 255, 0, 255})
	cpuInfText := newColoredLabel("CPU: 0.0", color.RGBA{0, 100, 100, 100})

	leftVBox := container.NewVBox(sysText, allocMemoryText, totalMemoryText, cpuInfText)
	leftContent := container.NewCenter(leftRectangle, leftVBox)
	go getSystemInformation(sysText, allocMemoryText, totalMemoryText, cpuInfText)

	rect := canvas.NewRectangle(nil)
	rect.SetMinSize(fyne.NewSize(800, 40))
	topHBox := container.NewHBox(Title)
	topContent := container.NewCenter(rect, topHBox)

	//right
	image := canvas.NewImageFromResource(data.FyneLogoTransparent)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(400, 600))
	rightRectangle := canvas.NewRectangle(color.RGBA{0, 0, 255, 255})
	rightRectangle.Resize(fyne.NewSize(400, 600))
	rightVBox := container.NewVBox(image)
	rightContent := container.NewCenter(rightRectangle, rightVBox)

	content := container.NewBorder(topContent, container.NewHBox(layout.NewSpacer(), status), leftContent, rightContent, nil)

	content.Resize(fyne.NewSize(800, 600))
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window, stopChan chan bool) *fyne.MainMenu {
	fileMenu := fyne.NewMenu("Run",
		fyne.NewMenuItem("Start", func() {
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
						go func() {
							err := runVerifier(RpcUrl.PlaceHolder, PrivateKey.Text, newWindow)
							if err != nil {
								dialog.ShowError(err, newWindow)
							}
						}()
					} else {
						go func() {
							err := runVerifier(RpcUrl.Text, PrivateKey.Text, newWindow)
							if err != nil {
								dialog.ShowError(err, newWindow)
							}
						}()
					}

					time.Sleep(2 * time.Second)
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
			go showLogs(logsTextArea, stopChan)
			newWindow.Resize(fyne.NewSize(600, 400))
			logsTextArea.Resize(fyne.NewSize(600, 400))
			newWindow.SetContent(logsTextArea)
			newWindow.Show()
			newWindow.SetOnClosed(func() {
				stopChan <- true
			})
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
	title.TextStyle = fyne.TextStyle{Bold: true, Italic: true, Monospace: true}
	title.Alignment = fyne.TextAlignCenter
}

func setTextArea(textArea *widget.Entry) {
	textArea.Resize(fyne.NewSize(1000, 800))
}

func showStatus(statusLabel *coloredLabel) {
	for {
		time.Sleep(2 * time.Second)
		data := checkVerifierIsActive()
		if data.Code == 200 {
			statusLabel.Text = "Active"
			statusLabel.Color = color.RGBA{0, 255, 0, 255}
		} else {
			statusLabel.Color = color.RGBA{255, 0, 0, 255}
		}

		statusLabel.Refresh()
	}
}

func showLogs(textArea *widget.Entry, stopChan <-chan bool) {
	for {
		select {
		case <-stopChan:
			return
		default:
			time.Sleep(2 * time.Second)
			textArea.SetText("")
			logsData := getSystemLogs()
			textArea.SetText(logsData.Result)
		}

	}
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
