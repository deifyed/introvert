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

	ui.statusbar.SetOnSubmitListener(func(address string) {
		go ui.Navigate(address)
	})

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

func (this *ui) Navigate(url string) {
	this.statusbar.SetAddress(url)

	page, err := html_utils.Parse(strings.NewReader(mockdata.MockRawWebpage))
	if err != nil {
		return
	}

	this.Open(asPage(page))
}

func (this *ui) Open(page page) {
	this.navbar.SetLinks(page.navigation)

	this.viewer.SetPageTitle(page.Title)
	this.viewer.SetSections(asViewerSections(page.Sections))
	this.viewer.Refresh()
}
