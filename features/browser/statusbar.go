package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type statusbar struct {
	container  *fyne.Container
	txtAddress *widget.Entry
}

func NewStatusBar() statusbar {
	frame := container.NewHBox()

	txtAddress := widget.NewEntry()

	frame.Add(txtAddress)

	return statusbar{
		container:  frame,
		txtAddress: txtAddress,
	}
}

func (this *statusbar) CanvasObject() fyne.CanvasObject {
	return this.container
}

func (this *statusbar) SetAddress(url string) {
	this.txtAddress.SetText(url)
}

func (this *statusbar) SetOnSubmitListener(fn func(string)) {
	this.txtAddress.OnSubmitted = fn
}
