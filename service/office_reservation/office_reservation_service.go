package office_reservation

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/rotem-ester/office-reservation-app/service/pkg/util"
)

type (
	OfficeReservationService struct {
		Reservations []OfficeReservation
	}

	OfficeReservationServiceHandlers interface {
		RevenueHandler(w http.ResponseWriter, r *http.Request)
		CapacityHandler(w http.ResponseWriter, r *http.Request)
	}
	OfficeReservationServiceOps interface {
		ParseData(data [][]string) error
		GetExpectedRevenueForMonth(year int, month time.Month) int
		GetExpectedCapacityForMonth(year int, month time.Month) int
	}
)

const DATE_LAYOUT = "2006-01-02"

func (ors *OfficeReservationService) RevenueHandler(w http.ResponseWriter, r *http.Request) {
	year, month, err := getParams(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res := ors.GetExpectedRevenueForMonth(year, month)

	fmt.Fprintf(w, "%v", res)
}

func (ors *OfficeReservationService) CapacityHandler(w http.ResponseWriter, r *http.Request) {
	year, month, err := getParams(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res := ors.GetExpectedCapacityForMonth(year, month)

	fmt.Fprintf(w, "%v", res)
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
					if val != "" {
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

func (ors *OfficeReservationService) GetExpectedRevenueForMonth(year int, month time.Month) int {
	var totalRevenue int

	for _, or := range ors.Reservations {
		totalRevenue += or.getMonthlyRevenue(year, month)
	}

	return totalRevenue
}

func (ors *OfficeReservationService) GetExpectedCapacityForMonth(year int, month time.Month) int {
	var totalCapacity int

	for _, or := range ors.Reservations {
		if or.getReservedDaysByMonth(year, month) == 0 {
			totalCapacity += or.Capacity
		}
	}

	return totalCapacity
}

func getParams(rawUrl *url.URL) (int, time.Month, error) {
	params, err := util.ParseQueryParams(rawUrl)
	if err != nil {
		return 0, 0, fmt.Errorf("failed parsing query params")
	}

	if err = util.EnsureParams(params); err != nil {
		return 0, 0, err
	}

	year, month, err := util.ParseParams(params)
	if err != nil {
		return 0, 0, err
	}

	return year, month, nil
}
