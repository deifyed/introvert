package browser

import (
	"fyne.io/fyne/v2/app"

	"github.com/deifyed/introvert/pkg/components/viewer"
)

func Start(url string) error {
	app := app.New()
	window := app.NewWindow("main")

	newUI(window)

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
