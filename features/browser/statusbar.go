package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type statusbar struct {
	container  *fyne.Container
	lblAddress *widget.Label
}

func NewStatusBar() statusbar {
	frame := container.NewHBox()

	lblAddress := widget.NewLabel("loading")
	frame.Add(lblAddress)

	return statusbar{
		container:  frame,
		lblAddress: lblAddress,
	}
}

func (this *statusbar) CanvasObject() fyne.CanvasObject {
	return this.container
}

func (this *statusbar) SetAddress(url string) {
	this.lblAddress.SetText(url)
}
