package datagrid

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var redrawParent func()
var uiScreen *tview.Table

func UI(redraw func(), data [][]string) *tview.Table {
	redrawParent = redraw

	table := tview.NewTable().
		SetBorders(true)

	cols, rows := len(data[0]), len(data)
	for row := range rows {
		for col := range cols {
			color := tcell.ColorWhite
			if row == 0 {
				color = tcell.ColorYellow
			}

			if col == 0 {
				if row == 0 {
					table.SetCell(row, 0,
						tview.NewTableCell(" ").
							SetTextColor(tcell.ColorYellow).
							SetAlign(tview.AlignRight))
				} else {
					table.SetCell(row, 0,
						tview.NewTableCell(strconv.Itoa(row)).
							SetTextColor(tcell.ColorYellow).
							SetAlign(tview.AlignRight))
				}
			}

			cell_value := data[row][col]
			_, parseNumberErr := strconv.ParseFloat(cell_value, 64)

			cellAlignment := tview.AlignLeft
			if parseNumberErr == nil {
				cellAlignment = tview.AlignCenter
			}

			table.SetCell(row, col+1,
				tview.NewTableCell(cell_value).
					SetTextColor(color).
					SetAlign(cellAlignment))
		}
	}
	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

	uiScreen = table

	return uiScreen
}
