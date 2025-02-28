package browse

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Start() error {
	app := app.New()

	window := app.NewWindow("Loading")
	window.SetContent(widget.NewLabel("Loading your webpage"))

	window.Show()

	app.Run()

	return nil
}
