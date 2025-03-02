package statusbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func New() Statusbar {
	frame := container.NewBorder(nil, nil, nil, nil)

	txtAddress := widget.NewEntry()
	txtAddress.Scroll = container.ScrollNone
	txtAddress.Wrapping = fyne.TextWrapOff

	txtAddress.SetPlaceHolder("Enter address")

	frame.Add(txtAddress)

	return Statusbar{
		container:  frame,
		txtAddress: txtAddress,
	}
}

func (this *Statusbar) CanvasObject() fyne.CanvasObject {
	return this.container
}

func (this *Statusbar) SetAddress(url string) {
	this.txtAddress.SetText(url)
}

func (this *Statusbar) SetOnSubmitListener(fn func(string)) {
	this.txtAddress.OnSubmitted = fn
}
