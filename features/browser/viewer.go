package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var size = fyne.NewSize(400, 0)

func NewViewer() viewer {
	c := container.NewMax()

	return viewer{
		container: c,
	}
}

type viewer struct {
	container *fyne.Container
}

func (this *viewer) CanvasObject() fyne.CanvasObject {
	return this.container
}

func (this *viewer) SetSections(sections []section) {
	this.container.RemoveAll()

	for _, s := range sections {
		this.container.Add(makeSection(s))
	}
}

func makeSection(s section) fyne.CanvasObject {
	wrapper := container.NewVBox()

	for _, p := range s.paragraphs {
		l := widget.NewMultiLineEntry()
		l.SetMinRowsVisible(5)
		l.SetText(p)
		l.Resize(size)

		wrapper.Add(l)
	}

	t := widget.NewCard(s.header, "", wrapper)

	return t
}
