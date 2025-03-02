package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type navigateListener func(string)

func NewNavbar(onNavigate navigateListener) navbar {
	return navbar{
		container:  container.NewVBox(),
		onNavigate: onNavigate,
	}
}

type navbar struct {
	container  *fyne.Container
	onNavigate navigateListener
}

func (this *navbar) CanvasObject() fyne.CanvasObject {
	return this.container
}

func (this *navbar) SetLinks(links []link) {
	this.container.RemoveAll()

	if len(links) == 0 {
		lblMissingNav := widget.NewLabel("No navigation found")

		this.container.Add(lblMissingNav)

		return
	}

	for _, l := range links {
		btn := widget.NewButton(l.title, func() {
			this.onNavigate(l.address)
		})

		this.container.Add(btn)
	}
}
