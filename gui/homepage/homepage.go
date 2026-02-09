package homepage

import (
	"math"
	"strconv"

	"github.com/janvdl/go-term-dataset-viewer/datacontroller"
	"github.com/janvdl/go-term-dataset-viewer/gui/datagrid"
	"github.com/rivo/tview"
)

var pages *tview.Pages
var pagesList *tview.List
var activePage string
var curr_idx int = 0

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
	totalPages := pages.GetPageCount()
	if curr_idx < totalPages-1 {
		curr_idx++
		changeActivePage(pages.GetPageNames(false)[curr_idx])
	}
}

func prevPage() {
	if curr_idx > 0 {
		curr_idx--
		changeActivePage(pages.GetPageNames(false)[curr_idx])
	}
}

func AddPage(data *datacontroller.Dataset) {
	if data != nil {
		pages.AddPage(data.Name, datagrid.UI(data.Data), true, true)
		refreshPagesList()
	}
}

func closeCurrentPage() {
	pages.RemovePage(activePage)

	// make sure curr page index is within bounds
	curr_idx = int(math.Max(float64(curr_idx), float64(pages.GetPageCount())))
	refreshPagesList()
}

func refreshPagesList() {
	pagesList.Clear()
	for idx, ds := range pages.GetPageNames(false) {
		active := activePage == ds
		if active {
			pagesList.AddItem(ds, strconv.Itoa(datacontroller.DatasetsLen[ds]-1)+" rows", ' ', nil).SetCurrentItem(idx)
		} else {
			pagesList.AddItem(ds, strconv.Itoa(datacontroller.DatasetsLen[ds]-1)+" rows", ' ', nil)
		}
	}
}

func Refresh() {
	refreshPagesList()
}
