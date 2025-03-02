package toolbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func New() Toolbar {
	frame := container.NewBorder(nil, nil, nil, nil)

	txtAddress := widget.NewEntry()
	txtAddress.Scroll = container.ScrollNone
	txtAddress.Wrapping = fyne.TextWrapOff

	txtAddress.SetPlaceHolder("Enter address")

	frame.Add(txtAddress)

	return Toolbar{
		container:  frame,
		txtAddress: txtAddress,
	}
}

func (this *Toolbar) CanvasObject() fyne.CanvasObject {
	return this.container
}

func (this *Toolbar) SetAddress(url string) {
	this.txtAddress.SetText(url)
}

func (this *Toolbar) SetOnSubmitListener(fn func(string)) {
	this.txtAddress.OnSubmitted = fn
}
