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
	window := app.NewWindow("main")

	setupUI(window)

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

func setupUI(window fyne.Window) {
	ui := &ui{
		statusbar: NewStatusBar(),
		viewport:  container.NewVBox(),
		navbar:    NewNavbar(),
	}

	ui.viewer = viewer.New(func() float32 {
		return window.Canvas().Size().Height - ui.statusbar.CanvasObject().Size().Height
	})

	ui.content = container.NewHSplit(ui.navbar.CanvasObject(), ui.viewer.CanvasObject())
	ui.content.SetOffset(.2)

	viewport := container.NewVBox(ui.statusbar.CanvasObject(), ui.content)

	go func() {
		ui.statusbar.SetAddress("mock URL")
		parsedPage, err := html_utils.Parse(strings.NewReader(mockdata.MockRawWebpage))
		if err != nil {
			return
		}

		page := asPage(parsedPage)

		ui.navbar.SetLinks(page.navigation)

		ui.viewer.SetPageTitle(page.Title)
		ui.viewer.SetSections(asViewerSections(page.Sections))
		ui.viewer.Refresh()
	}()

	window.SetContent(viewport)
	window.Show()
}

type ui struct {
	statusbar statusbar

	viewport *fyne.Container
	content  *container.Split

	navbar navbar
	viewer viewer.Viewer
}
