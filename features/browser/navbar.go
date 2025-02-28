package browser

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewNavbar() navbar {
	return navbar{
		container: container.NewVBox(),
	}
}

type navbar struct {
	container *fyne.Container
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
			navigate(l.address)
		})

		this.container.Add(btn)
	}
}

func navigate(url string) {
	fmt.Printf("Navigating to: %s\n", url)
}
