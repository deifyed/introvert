package browser

import (
	"fmt"
	"os"
	"strings"

	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//go:embed mockdata.html
var mockData string

func notify(app fyne.App, msg string) error {
	window := app.NewWindow("Notification")

	notification := widget.NewLabel(msg)
	window.SetContent(notification)

	return os.WriteFile("log.txt", []byte(msg), 0o777)
}

func Start(url string) error {
	app := app.New()

	navbar := container.New(layout.NewVBoxLayout())
	content := container.New(layout.NewVBoxLayout())

	window := app.NewWindow("main")
	window.SetContent(container.New(layout.NewHBoxLayout(), navbar, content))

	btnHome := widget.NewButton("Home", func() {
		navigate("/")
	})

	lblLoading := widget.NewLabel(fmt.Sprintf("Loading %s", url))

	navbar.Add(btnHome)
	content.Add(lblLoading)

	go func() {
		page, err := parse(strings.NewReader(mockData))
		if err != nil {
			lblLoading.Text = fmt.Sprintf("Error: %s", err.Error())
			lblLoading.Refresh()
			notify(app, err.Error())

			return
		}

		generateNav(navbar, page.navigation)

		lblLoading.Text = fmt.Sprintf("Page title: %s", page.Title)
		content.Refresh()
	}()

	window.ShowAndRun()

	return nil
}

func generateNav(container *fyne.Container, links []link) {
	for _, l := range links {
		btn := widget.NewButton(l.title, func() {
			navigate(l.address)
		})

		container.Add(btn)
	}
}

func navigate(url string) {
	fmt.Printf("Navigating to: %s\n", url)
}
