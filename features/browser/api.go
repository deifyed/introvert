package browser

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/deifyed/introvert/pkg/components/viewer"
	html_utils "github.com/deifyed/introvert/pkg/html"
	"github.com/deifyed/introvert/pkg/mockdata"
)

func Start(url string) error {
	app := app.New()

	// Setup containers
	window := app.NewWindow("main")

	setupUIOld(window)

	app.Run()

	return nil
}

func asViewerSections(originalSections []section) []viewer.Section {
	result := make([]viewer.Section, len(originalSections))

	for index, s := range originalSections {
		result[index] = viewer.Section{
			Title:      s.header,
			Paragraphs: s.paragraphs,
		}
	}

	return result
}

func setupUIOld(window fyne.Window) {
	viewport := container.NewVBox()
	main := container.NewHBox()

	// Setup parts
	navbar := NewNavbar()
	sb := NewStatusBar()

	viewer := viewer.New(func() float32 {
		return window.Canvas().Size().Height - sb.CanvasObject().Size().Height
	})

	// Bind main
	main.Add(navbar.CanvasObject())
	main.Add(viewer.CanvasObject())

	// Bind viewport
	viewport.Add(sb.CanvasObject())
	viewport.Add(main)

	// Bind window
	window.SetContent(viewport)

	go func() {
		sb.SetAddress("mock URL")
		parsedPage, err := html_utils.Parse(strings.NewReader(mockdata.MockRawWebpage))
		if err != nil {
			return
		}

		page := asPage(parsedPage)

		navbar.SetLinks(page.navigation)

		viewer.SetPageTitle(page.Title)
		viewer.SetSections(asViewerSections(page.Sections))
		viewer.Refresh()
	}()

	window.Show()
}
