package statusbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func New() Statusbar {
	frame := container.NewBorder(nil, nil, nil, nil)

	pbLoading := widget.NewProgressBarInfinite()

	frame.Add(pbLoading)

	return Statusbar{
		container: frame,
		pbLoading: pbLoading,
	}
}

func (this *Statusbar) CanvasObject() fyne.CanvasObject {
	return this.container
}

func (this *Statusbar) StartLoading() {
	this.pbLoading.Start()
	this.pbLoading.Hidden = false
	this.pbLoading.Refresh()
}

func (this *Statusbar) StopLoading() {
	this.pbLoading.Stop()
	this.pbLoading.Hidden = true
	this.pbLoading.Refresh()
}
