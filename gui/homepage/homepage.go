package homepage

import (
	"github.com/janvdl/go-term-dataset-viewer/drivers/csv"
	"github.com/janvdl/go-term-dataset-viewer/gui/datagrid"
	"github.com/rivo/tview"
)

var pages *tview.Pages
var pagesList *tview.List
var activePage string
var redrawParent func()

func UI(redraw func()) *tview.Grid {
	redrawParent = redraw

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	pagesList = tview.NewList()
	pages = tview.NewPages()

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(newPrimitive("Terminal Dataset Viewer"), 0, 0, 1, 2, 0, 0, false).
		AddItem(newPrimitive("V1.0 :: Jan van der Linde :: github.com/janvdl"), 2, 0, 1, 2, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(pagesList, 0, 0, 0, 0, 0, 0, false).
		AddItem(pages, 1, 0, 1, 2, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(pagesList, 1, 0, 1, 1, 0, 100, false).
		AddItem(pages, 1, 1, 1, 1, 0, 100, false)

	return grid
}

func ChangeActivePage(pageName string) {
	activePage = pageName
	pages.SwitchToPage(activePage)
	refreshPagesList()
}

func NextPage() {
	curr_idx := 0
	totalPages := pages.GetPageCount()
	for idx, page := range pages.GetPageNames(false) {
		if page == activePage {
			curr_idx = idx
		}
	}

	if curr_idx < totalPages-1 {
		curr_idx++
		ChangeActivePage(pages.GetPageNames(false)[curr_idx])
	}
}

func PrevPage() {
	curr_idx := 0
	for idx, page := range pages.GetPageNames(false) {
		if page == activePage {
			curr_idx = idx
		}
	}

	if curr_idx > 0 {
		curr_idx--
		ChangeActivePage(pages.GetPageNames(false)[curr_idx])
	}
}

func AddPage(pageName string) {
	data := csv.Read("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/pokemon.csv")
	pages.AddPage(pageName, datagrid.UI(redrawParent, data), true, true)
	ChangeActivePage(pageName)
	refreshPagesList()
}

func refreshPagesList() {
	pagesList.Clear()
	for _, page := range pages.GetPageNames(false) {
		secondaryText := ""
		if page == activePage {
			secondaryText = "currently viewing"
		}
		pagesList.AddItem(page, secondaryText, ' ', nil)
	}
}
