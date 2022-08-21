package file_util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LoadCsv(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open csv file. error: %w", err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: '%s'. error: %w", f.Name(), err)
	}

	return data, nil
}