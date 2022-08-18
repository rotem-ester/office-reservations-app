package util

import (
	"time"
)

func GetDaysNumByYearAndMonth(year int, month time.Month) int {
	t := time.Date(year, month, 1, 23, 0, 0, 0, time.UTC)
	t = t.AddDate(0, 1, -1)
	return t.Day()
}