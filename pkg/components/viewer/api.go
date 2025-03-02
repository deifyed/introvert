package viewer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Viewer struct {
	getAvailableHeight heightGetter

	container *fyne.Container
	scroll    *container.Scroll

	lblPageTitle *widget.Label

	title    string
	sections []Section
}

func New(getAvailableHeight heightGetter) Viewer {
	c := container.NewVBox()

	lblPageTitle := widget.NewLabelWithStyle("", fyne.TextAlignLeading, pageHeaderStyle)
	c.Add(lblPageTitle)

	return Viewer{
		container:          c,
		scroll:             container.NewScroll(c),
		getAvailableHeight: getAvailableHeight,

		lblPageTitle: lblPageTitle,

		title:    "",
		sections: []Section{},
	}
}

func (this *Viewer) Refresh() {
	this.container.RemoveAll()

	this.lblPageTitle.SetText(this.title)
	this.container.Add(this.lblPageTitle)

	for _, s := range this.sections {
		this.container.Add(makeSection(s))
	}

	this.refreshSize()
}

func (this *Viewer) CanvasObject() fyne.CanvasObject {
	return this.scroll
}

func (this *Viewer) SetPageTitle(title string) {
	this.title = title
}

func (this *Viewer) SetSections(sections []Section) {
	this.sections = sections
}
