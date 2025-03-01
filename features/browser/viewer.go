package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type heightGetter func() float32

func NewViewer(getAvailableHeight heightGetter) viewer {
	c := container.NewVBox()
	scroll := container.NewScroll(c)

	return viewer{
		container:          c,
		scroll:             scroll,
		getAvailableHeight: getAvailableHeight,
	}
}

type viewer struct {
	container          *fyne.Container
	scroll             *container.Scroll
	getAvailableHeight heightGetter
}

func (this *viewer) refreshSize() {
	size := fyne.NewSize(500, this.getAvailableHeight())

	this.scroll.SetMinSize(size)
}

func (this *viewer) CanvasObject() fyne.CanvasObject {
	return this.scroll
}

func (this *viewer) SetSections(sections []section) {
	this.container.RemoveAll()

	for _, s := range sections {
		this.container.Add(makeSection(s))
	}

	this.refreshSize()
}

func makeSection(s section) fyne.CanvasObject {
	wrapper := container.NewVBox()

	for _, p := range s.paragraphs {
		l := widget.NewLabel(p)
		l.Wrapping = fyne.TextWrapWord

		wrapper.Add(l)
	}

	t := widget.NewCard(s.header, "", wrapper)

	return t
}
