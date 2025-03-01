package browser

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	html_utils "github.com/deifyed/introvert/pkg/html"
	"github.com/deifyed/introvert/pkg/mockdata"
)

func Start(url string) error {
	app := app.New()

	// Setup containers
	window := app.NewWindow("main")
	viewport := container.NewVBox()
	main := container.NewHBox()

	// Setup parts
	navbar := NewNavbar()
	sb := NewStatusBar()

	viewer := NewViewer(func() float32 {
		return window.Canvas().Size().Height - sb.CanvasObject().Size().Height
	})

	// Bind main
	main.Add(navbar.CanvasObject())
	main.Add(viewer.CanvasObject())

	// Bind viewport
	viewport.Add(sb.CanvasObject())
	viewport.Add(main)

	window.SetContent(viewport)

	go func() {
		sb.SetAddress(url)
		parsedPage, err := html_utils.Parse(strings.NewReader(mockdata.MockRawWebpage))
		if err != nil {
			notify(app, err.Error())

			return
		}

		fmt.Println("no error")

		page := asPage(parsedPage)

		window.SetTitle(page.Title)
		navbar.SetLinks(page.navigation)
		viewer.SetSections(page.Sections)
	}()

	window.Show()
	app.Run()

	return nil
}
