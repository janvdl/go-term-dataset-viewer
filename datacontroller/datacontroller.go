package datacontroller

import (
	"path/filepath"
	"strings"

	"github.com/janvdl/go-term-dataset-viewer/drivers/csv"
	"github.com/janvdl/go-term-dataset-viewer/drivers/xpt"
)

var Datasets []*Dataset

type Dataset struct {
	Name string
	Data [][]string
}

func OpenData(filePath string) {
	ds := &Dataset{}

	ext := strings.ToLower(filepath.Ext(filePath))
	ds.Name = filepath.Base(filePath)

	if ext == ".csv" {
		data, err := csv.Read(filePath)

		if err != nil {
			panic(err)
		}

		ds.Data = data
		Datasets = append(Datasets, ds)
	} else if ext == ".xpt" {
		data, err := xpt.Read(filePath)

		if err != nil {
			panic(err)
		}

		ds.Data = data
		Datasets = append(Datasets, ds)
	}
}
