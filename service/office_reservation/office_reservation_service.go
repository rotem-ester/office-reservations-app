package office_reservation

import (
	"fmt"
	"net/http"
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

func (ors *OfficeReservationService) RevenueHandler(w http.ResponseWriter, r *http.Request) {
	params, err := util.ParseQueryParams(r.URL)
	if err != nil {
		fmt.Printf("failed parsing revenue query params: %s", err.Error())
		http.Error(w, "failed parsing query params", http.StatusInternalServerError)
	}

	if err = util.EnsureRevenueParams(params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	year, month, err := util.ParseRevenueParams(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res := ors.GetExpectedRevenueForMonth(year, month)
	
	fmt.Fprintf(w, "%v", res)
}

func (ors *OfficeReservationService) CapacityHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "capacity handler")
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