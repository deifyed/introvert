package browser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func notify(app fyne.App, msg string) {
	window := app.NewWindow("Notification")

	notification := widget.NewLabel(msg)
	notification.Wrapping = fyne.TextWrapWord
	window.SetContent(notification)

	window.Show()

	return
}
