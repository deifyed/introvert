package browser

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/deifyed/introvert/pkg/components/toolbar"
	"github.com/deifyed/introvert/pkg/components/viewer"
	html_utils "github.com/deifyed/introvert/pkg/html"
	"github.com/deifyed/introvert/pkg/mockdata"
)

const (
	viewEmpty = iota
	viewLoading
	viewContent
)

type ui struct {
	toolbar toolbar.Toolbar

	viewport *fyne.Container
	content  *fyne.Container

	navbar navbar
	viewer viewer.Viewer

	currentView int
}

func (this *ui) Open(page page) {
	this.navbar.SetLinks(page.navigation)

	this.viewer.SetPageTitle(page.Title)
	this.viewer.SetSections(asViewerSections(page.Sections))
	this.viewer.Refresh()
}

func newUI(window fyne.Window, url string) *ui {
	ui := &ui{
		toolbar: toolbar.New(),
		content: container.NewBorder(nil, nil, nil, nil),
		navbar:  NewNavbar(),
	}

	ui.navbar = NewNavbar(ui.Navigate)

	ui.viewer = viewer.New(func() float32 {
		return window.Canvas().Size().Height - ui.toolbar.CanvasObject().Size().Height
	})

	ui.viewport = container.NewVBox(ui.toolbar.CanvasObject(), ui.content)

	ui.toolbar.SetOnSubmitListener(func(address string) {
		go ui.Navigate(ensureURL(address))
	})

	ui.showEmptyScreen()

	window.SetContent(ui.viewport)
	window.Show()

	if url != "" {
		go ui.Navigate(url)
	}

	return ui
}

func (this *ui) Navigate(url string) {
	this.toolbar.SetAddress(url)

	this.showLoading()

	page, err := html_utils.Parse(strings.NewReader(mockdata.MockRawWebpage))
	if err != nil {
		return
	}

	time.Sleep(2 * time.Second)

	this.showContent()

	this.Open(asPage(page))
}

func (this *ui) showContent() {
	this.content.RemoveAll()

	splitter := container.NewHSplit(this.navbar.CanvasObject(), this.viewer.CanvasObject())
	splitter.SetOffset(.2)

	this.content.Add(splitter)

	this.currentView = viewContent
}

func (this *ui) showLoading() {
	this.content.RemoveAll()

	pbLoading := widget.NewProgressBarInfinite()
	pbLoading.Start()

	this.content.Add(pbLoading)

	this.currentView = viewLoading
}

func (this *ui) showEmptyScreen() {
	this.content.RemoveAll()

	label := widget.NewLabel("⭐")
	label.Alignment = fyne.TextAlignCenter

	this.content.Add(label)

	this.currentView = viewEmpty
}

func ensureURL(rawURL string) string {
	if len(strings.Split(rawURL, "https://")) == 1 {
		return fmt.Sprintf("https://%s", rawURL)
	}

	return rawURL
}
