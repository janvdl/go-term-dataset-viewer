package mainmenu

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func LoadApp(data [][]string) {
	app := tview.NewApplication()
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
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
