package main

import (
	"github.com/janvdl/go-term-dataset-viewer/drivers/csv"
	"github.com/janvdl/go-term-dataset-viewer/ui/mainmenu"
)

func main() {
	data := csv.ReadCsv("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/pokemon.csv")
	mainmenu.LoadApp(data)
}
