package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/janvdl/go-term-dataset-viewer/datacontroller"
	"github.com/janvdl/go-term-dataset-viewer/gui/homepage"
	"github.com/rivo/tview"
)

var app *tview.Application

//var home *tview.Grid

func Init() {
	app = tview.NewApplication()

	home := homepage.UI(app)

	files := []string{
		"/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/bulbasaur.csv",
		"/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/sample.xpt",
	}

	for idx := range files {
		ds := datacontroller.OpenData(files[idx])
		homepage.AddPage(ds)
	}

	homepage.Refresh()

	app.SetInputCapture(keyPressHandler)
	if err := app.SetRoot(home, true).SetFocus(home).Run(); err != nil {
		panic(err)
	}
}

func keyPressHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Rune() == 'q' {
		app.Stop()
		return nil
	} else {
		homepage.KeyPressHandler(eventKey)
	}

	return eventKey
}
