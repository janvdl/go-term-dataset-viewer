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

	datacontroller.OpenData("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/bulbasaur.csv")
	datacontroller.OpenData("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/sample.xpt")

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
