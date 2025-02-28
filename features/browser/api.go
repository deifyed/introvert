package browser

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Start(url string) error {
	app := app.New()

	viewport := container.New(layout.NewVBoxLayout())

	statusbar := NewStatusBar()
	statusbar.SetAddress(url)
	viewport.Add(statusbar.CanvasObject())

	main := container.NewHBox()

	navbar := NewNavbar()
	main.Add(navbar.CanvasObject())

	viewer := NewViewer()
	main.Add(viewer.CanvasObject())

	viewport.Add(main)

	window := app.NewWindow("main")
	window.SetContent(viewport)

	go func() {
		page, err := parse(strings.NewReader(mockData))
		if err != nil {
			app.SendNotification(fyne.NewNotification("error", err.Error()))

			return
		}

		window.SetTitle(page.Title)
		navbar.SetLinks(page.navigation)
		viewer.SetSections(page.Sections)
	}()

	window.ShowAndRun()

	return nil
}
