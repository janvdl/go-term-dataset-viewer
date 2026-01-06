package main

import (
	"github.com/janvdl/go-term-dataset-viewer/drivers/csv"
	"github.com/janvdl/go-term-dataset-viewer/gui/datagrid"
)

func main() {
	data := csv.Read("/Users/jvdl/Programming/github.com/janvdl/go-term-dataset-viewer/sampledata/pokemon.csv")
	datagrid.LoadApp(data)
}
