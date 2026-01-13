package homepage

import (
	"github.com/gdamore/tcell/v2"
)

func KeyPressHandler(eventKey *tcell.EventKey) *tcell.EventKey {
	if eventKey.Rune() == '[' {
		prevPage()
	} else if eventKey.Rune() == ']' {
		nextPage()
	} else if eventKey.Rune() == 'x' {
		closeCurrentPage()
	}

	return eventKey
}
