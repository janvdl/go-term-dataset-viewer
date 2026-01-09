package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/janvdl/go-term-dataset-viewer/gui/homepage"
	"github.com/rivo/tview"
)

var app *tview.Application

//var home *tview.Grid

func Init() {
	app = tview.NewApplication()

	redraw := func() {
		app.Draw()
	}

	home := homepage.UI(redraw)
	homepage.AddPage("pokemon1")
	homepage.AddPage("pokemon2")
	homepage.AddPage("pokemon3")
	homepage.ChangeActivePage("pokemon2")

	app.SetInputCapture(keyPressHandler)
	if err := app.SetRoot(home, true).SetFocus(home).Run(); err != nil {
		panic(err)
	}
}

func keyPressHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Rune() == 'q' {
		app.Stop()
		return nil
	} else if eventKey.Rune() == '[' {
		homepage.PrevPage()
	} else if eventKey.Rune() == ']' {
		homepage.NextPage()
	}

	return eventKey
}
