package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type heightGetter func() float32

var titleStyle = fyne.TextStyle{
	Bold:      true,
	Italic:    false,
	Monospace: false,
	Symbol:    false,
	TabWidth:  0,
	Underline: false,
}

func NewViewer(getAvailableHeight heightGetter) viewer {
	c := container.NewVBox()

	lblPageTitle := widget.NewLabelWithStyle("", fyne.TextAlignLeading, titleStyle)
	c.Add(lblPageTitle)

	scroll := container.NewScroll(c)

	return viewer{
		container:          c,
		scroll:             scroll,
		getAvailableHeight: getAvailableHeight,

		lblPageTitle: lblPageTitle,

		title:    "",
		sections: []section{},
	}
}

type viewer struct {
	getAvailableHeight heightGetter

	container *fyne.Container
	scroll    *container.Scroll

	lblPageTitle *widget.Label

	title    string
	sections []section
}

func (this *viewer) refreshSize() {
	height := this.getAvailableHeight() - this.lblPageTitle.Size().Height
	var width float32 = 500

	this.scroll.SetMinSize(fyne.NewSize(width, height))
}

func (this *viewer) Refresh() {
	this.container.RemoveAll()

	this.lblPageTitle.SetText(this.title)
	this.container.Add(this.lblPageTitle)

	for _, s := range this.sections {
		this.container.Add(makeSection(s))
	}

	this.refreshSize()
}

func (this *viewer) CanvasObject() fyne.CanvasObject {
	return this.scroll
}

func (this *viewer) SetPageTitle(title string) {
	this.title = title
}

func (this *viewer) SetSections(sections []section) {
	this.sections = sections
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
