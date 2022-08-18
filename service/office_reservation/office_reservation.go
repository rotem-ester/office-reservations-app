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
	
	return 0
}

func (or *OfficeReservation) getReservedDaysPerMonth(year int, month time.Month) int {
	return 0
}

func (or *OfficeReservation) getDailyPriceByMonth(year int, month time.Month) int {
	days := util.GetDaysNumByYearAndMonth(year, month)
	return or.MonthlyPrice / days
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


