package gui

import "github.com/rivo/tview"

var app *tview.Application
var pages *tview.Pages
var activePage int

func Init() {
	app = tview.NewApplication()
	pages = tview.NewPages()

	redraw := func() {
		app.Draw()
	}
}
