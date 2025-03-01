package viewer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Section struct {
	Title      string
	Paragraphs []string
}

func makeSection(s Section) fyne.CanvasObject {
	wrapper := container.NewVBox()

	for _, p := range s.Paragraphs {
		l := widget.NewLabel(p)
		l.Wrapping = fyne.TextWrapWord

		wrapper.Add(l)
	}

	t := widget.NewCard(s.Title, "", wrapper)

	return t
}
