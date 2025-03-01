package browser

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	html_utils "github.com/deifyed/introvert/pkg/html"
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
		parsedPage, err := html_utils.Parse(strings.NewReader(mockData))
		if err != nil {
			notify(app, err.Error())

			return
		}

		fmt.Println("no error")

		page := asPage(parsedPage)

		window.SetTitle(page.Title)
		navbar.SetLinks(page.navigation)
		viewer.SetSections(page.Sections)
		viewerScroll.SetMinSize(calculateViewerSize(window, &sb))
	}()

	window.Show()
	app.Run()

	return nil
}

func calculateViewerSize(window fyne.Window, sb *statusbar) fyne.Size {
	var width float32 = 500
	var height float32 = window.Canvas().Size().Height - sb.CanvasObject().Size().Height

	return fyne.NewSize(width, height)
}
