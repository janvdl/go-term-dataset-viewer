package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/janvdl/go-term-dataset-viewer/drivers/csv"
	"github.com/janvdl/go-term-dataset-viewer/drivers/xpt"
	"github.com/janvdl/go-term-dataset-viewer/gui/homepage"
	"github.com/rivo/tview"
)

var app *tview.Application

//var home *tview.Grid

func Init() {
	app = tview.NewApplication()

	home := homepage.UI(app)

	data_bulbasaur, _ := csv.Read("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/bulbasaur.csv")
	homepage.AddPage("bulbasaur", data_bulbasaur)

	data_charmander, _ := csv.Read("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/charmander.csv")
	homepage.AddPage("charmander", data_charmander)

	data_squirtle, _ := csv.Read("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/squirtle.csv")
	homepage.AddPage("squirtle", data_squirtle)

	data_xpt, _ := xpt.Read("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/sample.xpt")
	homepage.AddPage("xpt example", data_xpt)

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
