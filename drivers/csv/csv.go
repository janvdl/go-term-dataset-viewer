package csv

import (
	"encoding/csv"
	"os"
)

func Read(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return data
}
