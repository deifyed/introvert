package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewViewer() viewer {
	c := container.NewVBox()

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
		l := widget.NewLabel(p)
		l.Wrapping = fyne.TextWrapWord

		wrapper.Add(l)
	}

	t := widget.NewCard(s.header, "", wrapper)

	return t
}
