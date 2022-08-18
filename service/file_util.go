package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const DATE_LAYOUT = "2006-01-02"

func LoadCsv(path string) ([]OfficeReservation, error) {
	var reservations []OfficeReservation
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open csv file '%s'. error: %w", path, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: '%s'. error: %w", path, err)
	}

	for i, line := range data {
		if i > 0 {
			var or OfficeReservation

			for j, val := range line {
				if j == 0 {
					or.Capacity, err = strconv.Atoi(strings.TrimSpace(val))
					if err != nil {
						return nil, fmt.Errorf("failed to parse value from csv file: '%s'. error: %w", path, err)
					}
				}

				if j == 1 {
					or.MonthlyPrice, err = strconv.Atoi(strings.TrimSpace(val))
					if err != nil {
						return nil, fmt.Errorf("failed to parse value from csv file: '%s'. error: %w", path, err)
					}
				}

				if j == 2 {
					or.StartDay, err = time.Parse(DATE_LAYOUT, strings.TrimSpace(val))
					if err != nil {
						return nil, fmt.Errorf("failed to parse value from csv file: '%s'. error: %w", path, err)
					}
				}

				if j == 3 {
					if val != ""{
						or.EndDay, err = time.Parse(DATE_LAYOUT, strings.TrimSpace(val))
						if err != nil {
							return nil, fmt.Errorf("failed to parse value from csv file: '%s'. error: %w", path, err)
						}
					}
				}
			}
			reservations = append(reservations, or)
		}
	}

	return reservations, nil
}