package viewer

import "fyne.io/fyne/v2"

type heightGetter func() float32

func (this *Viewer) refreshSize() {
	height := this.getAvailableHeight() - this.lblPageTitle.Size().Height
	var width float32 = 500

	this.scroll.SetMinSize(fyne.NewSize(width, height))
}
