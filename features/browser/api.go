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

	sb := NewStatusBar()
	sb.SetAddress(url)
	viewport.Add(sb.CanvasObject())

	main := container.NewHBox()

	navbar := NewNavbar()
	main.Add(navbar.CanvasObject())

	viewer := NewViewer()

	viewerScroll := container.NewScroll(viewer.CanvasObject())

	main.Add(viewerScroll)

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
		viewerScroll.SetMinSize(calculateViewerSize(window, &sb))
	}()

	window.ShowAndRun()

	return nil
}

func calculateViewerSize(window fyne.Window, sb *statusbar) fyne.Size {
	var width float32 = 500
	var height float32 = window.Canvas().Size().Height - sb.CanvasObject().Size().Height

	return fyne.NewSize(width, height)
}
