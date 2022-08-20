package office_reservation

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	util "github.com/rotem-ester/office-reservation-app/service/pkg/util"
)

const DATE_LAYOUT = "2006-01-02"

type (
	OfficeReservationService struct {
		Reservations []OfficeReservation
	}

	OfficeReservationServiceOps interface {
		ParseData(data [][]string) error
	}

	OfficeReservation struct {
		Capacity     int
		MonthlyPrice int
		StartDay     time.Time
		EndDay       time.Time
	}

	OfficeReservationOps interface {
		GetMonthlyRevenue(year int, month time.Month) int
	}
)

func (or *OfficeReservation) GetMonthlyRevenue(year int, month time.Month) int {
	days := float64(or.getReservedDaysByMonth(year, month))
	dailyPrice := or.getDailyPriceByMonth(year,month)

	return int(days * dailyPrice)
}

func (or *OfficeReservation) getReservedDaysByMonth(year int, month time.Month) int {
	tempStartDay := or.StartDay
	tempEndDay := or.EndDay
	monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	monthEndDay := util.GetDaysNumByYearAndMonth(year, month)
	monthEnd := time.Date(year, month, monthEndDay, 0, 0, 0, 0, time.UTC)

	if !or.EndDay.IsZero() && or.EndDay.Before(monthStart) {
		return 0
	}

	if or.StartDay.After(monthEnd) {
		return 0
	}

	if or.StartDay.Before(monthStart) {
		tempStartDay = monthStart
	}

	if or.EndDay.After(monthEnd) || or.EndDay.IsZero() {
		tempEndDay = monthEnd
	}
	
	tempStartDay = tempStartDay.AddDate(0, 0, -1) // we want to incluse the start day in the calculation
	difference := tempEndDay.Sub(tempStartDay)
	return int(difference.Hours() / 24)
}

func (or *OfficeReservation) getDailyPriceByMonth(year int, month time.Month) float64 {
	days := float64(util.GetDaysNumByYearAndMonth(year, month))
	price := float64(or.MonthlyPrice) / days
	fmt.Printf("days: %v\n", days)
	fmt.Printf("price: %v\n", price)
	return price
}

func (ors *OfficeReservationService) ParseData(data [][]string) error {
	var reservations []OfficeReservation
	var err error
	
	for i, line := range data {
		if i > 0 {
			var or OfficeReservation

			for j, val := range line {
				if j == 0 {
					or.Capacity, err = strconv.Atoi(strings.TrimSpace(val))
					if err != nil {
						return fmt.Errorf("failed to parse value from data. error: %w", err)
					}
				}

				if j == 1 {
					or.MonthlyPrice, err = strconv.Atoi(strings.TrimSpace(val))
					if err != nil {
						return fmt.Errorf("failed to parse value from data. error: %w", err)
					}
				}

				if j == 2 {
					or.StartDay, err = time.Parse(DATE_LAYOUT, strings.TrimSpace(val))
					if err != nil {
						return fmt.Errorf("failed to parse value from data. error: %w", err)
					}
				}

				if j == 3 {
					if val != ""{
						or.EndDay, err = time.Parse(DATE_LAYOUT, strings.TrimSpace(val))
						if err != nil {
							return fmt.Errorf("failed to parse value from data. error: %w", err)
						}
					}
				}
			}
			reservations = append(reservations, or)
		}
	}

	ors.Reservations = reservations
	return nil
}


