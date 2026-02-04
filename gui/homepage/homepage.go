package homepage

import (
	"github.com/janvdl/go-term-dataset-viewer/datacontroller"
	"github.com/janvdl/go-term-dataset-viewer/gui/datagrid"
	"github.com/rivo/tview"
)

var pages *tview.Pages
var pagesList *tview.List
var activePage string

func UI(app *tview.Application) *tview.Grid {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	pagesList = tview.NewList()
	pages = tview.NewPages()

	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(newPrimitive("Terminal Dataset Viewer"), 0, 0, 1, 2, 0, 0, false).
		AddItem(newPrimitive("V0.1 :: Jan van der Linde :: github.com/janvdl"), 2, 0, 1, 2, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(newPrimitive("Datasets"), 1, 0, 1, 1, 0, 0, false).
		AddItem(pagesList, 1, 0, 1, 1, 0, 0, false).
		AddItem(pages, 1, 1, 1, 1, 0, 0, true)

	app.SetFocus(pages)

	return grid
}

func changeActivePage(pageName string) {
	activePage = pageName
	pages.SwitchToPage(activePage)
	refreshPagesList()
}

func nextPage() {
	curr_idx := 0
	totalPages := pages.GetPageCount()
	for idx, page := range pages.GetPageNames(false) {
		if page == activePage {
			curr_idx = idx
		}
	}

	if curr_idx < totalPages-1 {
		curr_idx++
		changeActivePage(pages.GetPageNames(false)[curr_idx])
	}
}

func prevPage() {
	curr_idx := 0
	for idx, page := range pages.GetPageNames(false) {
		if page == activePage {
			curr_idx = idx
		}
	}

	if curr_idx > 0 {
		curr_idx--
		changeActivePage(pages.GetPageNames(false)[curr_idx])
	}
}

func AddPage(pageName string, data *datacontroller.Dataset) {
	if data != nil {
		pages.AddPage(pageName, datagrid.UI(data.Data), true, true)
		changeActivePage(pageName)
		refreshPagesList()
	}
}

func closeCurrentPage() {
	pages.RemovePage(activePage)
	refreshPagesList()
}

func refreshPagesList() {
	pagesList.Clear()
	for _, ds := range datacontroller.Datasets {
		pagesList.AddItem(ds.Name, "secondary", ' ', nil)
	}
}

func Refresh() {
	refreshPagesList()
}
