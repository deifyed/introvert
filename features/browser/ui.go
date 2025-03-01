package browser

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/deifyed/introvert/pkg/components/statusbar"
	"github.com/deifyed/introvert/pkg/components/viewer"
	html_utils "github.com/deifyed/introvert/pkg/html"
	"github.com/deifyed/introvert/pkg/mockdata"
)

type ui struct {
	statusbar statusbar.Statusbar

	viewport *fyne.Container
	content  *container.Split

	navbar navbar
	viewer viewer.Viewer
}

func (this *ui) Open(page page) {
	this.navbar.SetLinks(page.navigation)

	this.viewer.SetPageTitle(page.Title)
	this.viewer.SetSections(asViewerSections(page.Sections))
	this.viewer.Refresh()
}

func newUI(window fyne.Window) *ui {
	ui := &ui{
		statusbar: statusbar.New(),
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

	return ui
}

func (this *ui) Navigate(url string) {
	this.statusbar.SetAddress(url)

	page, err := html_utils.Parse(strings.NewReader(mockdata.MockRawWebpage))
	if err != nil {
		return
	}

	this.Open(asPage(page))
}
