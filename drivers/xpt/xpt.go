package xpt

import (
	"os"

	goxpt "github.com/janvdl/go-xpt"
)

func Read(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := goxpt.ReadXPT(filePath)
	if err != nil {
		panic(err)
	}

	return data.AsSimpleGrid(), nil
}
