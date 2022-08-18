package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LoadCsv(file *os.File) ([][]string, error) {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: '%s'. error: %w", file.Name(), err)
	}

	return data, nil
}