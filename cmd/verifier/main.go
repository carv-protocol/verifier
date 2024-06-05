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
	"github.com/carv-protocol/verifier/internal/common"
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
	// verifier status
	status := newColoredLabel("InActive", color.RGBA{255, 0, 0, 255})
	go showStatus(status)

	// top
	image := canvas.NewImageFromResource(data.FyneLogoTransparent)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(100, 40))
	rect := canvas.NewRectangle(nil)
	rect.SetMinSize(fyne.NewSize(800, 40))
	topContent := container.NewBorder(container.NewHBox(image), nil, nil, nil, nil)

	// left
	leftRectangle := canvas.NewRectangle(nil)
	leftRectangle.SetMinSize(fyne.NewSize(500, 600))
	leftRectangle.FillColor = color.RGBA{R: 0, A: 255}

	sysText := newColoredLabel("System", color.White)
	sysTextRectangle := canvas.NewRectangle(nil)
	sysTextRectangle.SetMinSize(fyne.NewSize(400, 40))
	sysTextRectangle.FillColor = color.RGBA{R: 0, A: 200}
	sysTextRectangleContent := container.NewCenter(sysTextRectangle, sysText)

	allocMemoryText := newColoredLabel("Allocated Memory: 0.0", color.White)
	allocTextRectangle := canvas.NewRectangle(nil)
	allocTextRectangle.SetMinSize(fyne.NewSize(400, 40))
	allocTextRectangle.FillColor = color.RGBA{R: 255, A: 200}
	allocTextRectangleContent := container.NewCenter(allocTextRectangle, allocMemoryText)

	totalMemoryText := newColoredLabel("Total Memory: 0.0", color.White)
	totalMemoryTextRectangle := canvas.NewRectangle(nil)
	totalMemoryTextRectangle.SetMinSize(fyne.NewSize(400, 40))
	totalMemoryTextRectangle.FillColor = color.RGBA{B: 255, A: 200}
	totalMemoryTextRectangleContent := container.NewCenter(totalMemoryTextRectangle, totalMemoryText)

	cpuInfText := newColoredLabel("CPU: 0.0", color.White)
	cpuInfTextRectangle := canvas.NewRectangle(nil)
	cpuInfTextRectangle.SetMinSize(fyne.NewSize(400, 40))
	cpuInfTextRectangle.FillColor = color.RGBA{B: 100, A: 200}
	cpuInfTextRectangleContent := container.NewCenter(cpuInfTextRectangle, cpuInfText)

	leftContent := container.NewCenter(leftRectangle, container.NewVBox(sysTextRectangleContent, allocTextRectangleContent, totalMemoryTextRectangleContent, cpuInfTextRectangleContent))
	go getSystemInformation(sysText, allocMemoryText, totalMemoryText, cpuInfText)

	//right
	rightRectangle := canvas.NewRectangle(nil)
	rightRectangle.SetMinSize(fyne.NewSize(500, 600))
	rightRectangle.FillColor = color.White
	// current block
	currentBlockText := newColoredLabel("Current Block Height: 0", color.White)
	currentBlockTextRectangle := canvas.NewRectangle(nil)
	currentBlockTextRectangle.SetMinSize(fyne.NewSize(400, 40))
	currentBlockTextRectangle.FillColor = color.RGBA{R: 0, A: 200}
	currentBlockTextRectangleContent := container.NewCenter(currentBlockTextRectangle, currentBlockText)
	go logWatcher(currentBlockText)

	latestBlockText := newColoredLabel("Latest Block Height: 0", color.White)
	latestBlockTextRectangle := canvas.NewRectangle(nil)
	latestBlockTextRectangle.SetMinSize(fyne.NewSize(400, 40))
	latestBlockTextRectangle.FillColor = color.RGBA{R: 100, A: 200}
	latestBlockTextRectangleContent := container.NewCenter(latestBlockTextRectangle, latestBlockText)

	go getLatestBlockFromChain(latestBlockText)

	rightRectangleContent := container.NewCenter(rightRectangle, container.NewVBox(currentBlockTextRectangleContent, latestBlockTextRectangleContent))
	go getSystemInformation(sysText, allocMemoryText, totalMemoryText, cpuInfText)

	//middle

	// bottom
	bottomHBox := container.NewHBox(layout.NewSpacer(), status)
	content := container.NewBorder(topContent, bottomHBox, leftContent, rightRectangleContent, nil)

	// filter log
	myWindow.SetContent(content)
	myWindow.FullScreen()

	myWindow.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window, stopChan chan bool) *fyne.MainMenu {
	fileMenu := fyne.NewMenu("Run",
		fyne.NewMenuItem("Start", func() {
			newWindow := a.NewWindow("RunNode")
			// form
			RpcUrl := widget.NewEntry()
			RpcUrl.SetPlaceHolder(common.OP_BNB_RPC_URL)
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
	type Transaction struct {
		ID     string
		Amount string
		TxHash string
		Date   string
	}

	//transactionMenu := fyne.NewMenu("Transaction",
	//	fyne.NewMenuItem("History", func() {
	//		newWindow := a.NewWindow("TransactionWindow")
	//		// get transaction history
	//		transactions := []Transaction{
	//			{"1", "100", "0x0001", "2021-09-01"},
	//			{"2", "200", "0x0002", "2021-09-02"},
	//			{"3", "300", "0x0003", "2021-09-03"},
	//		}
	//
	//		table := widget.NewTable(
	//			func() (int, int) {
	//				return len(transactions) + 1, 4
	//			},
	//			func() fyne.CanvasObject {
	//				return widget.NewLabel("")
	//			},
	//			func(id widget.TableCellID, cell fyne.CanvasObject) {
	//				label := cell.(*widget.Label)
	//
	//				if id.Row == 0 {
	//					switch id.Col {
	//					case 0:
	//						label.SetText("ID")
	//					case 1:
	//						label.SetText("Amount")
	//					case 2:
	//						label.SetText("TxHash")
	//					case 3:
	//						label.SetText("Date")
	//					}
	//				} else {
	//					switch id.Col {
	//					case 0:
	//						label.SetText(transactions[id.Row-1].ID)
	//					case 1:
	//						label.SetText(transactions[id.Row-1].Amount)
	//					case 2:
	//						label.SetText(transactions[id.Row-1].TxHash)
	//					case 3:
	//						label.SetText(transactions[id.Row-1].Date)
	//					}
	//				}
	//			},
	//		)
	//		table.SetColumnWidth(0, 100)
	//		table.SetColumnWidth(1, 200)
	//		table.SetColumnWidth(2, 300)
	//		table.SetColumnWidth(3, 300)
	//		newWindow.SetContent(table)
	//		newWindow.Resize(fyne.NewSize(600, 400))
	//		newWindow.Show()
	//	}),
	//)
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
