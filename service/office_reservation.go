package main

import "time"

type OfficeReservationOps interface {
	getReservedDaysPerMonth(year int, month time.Month) int
	getDailyPriceByMonth(year int, month time.Month) int
	getMonthlyRevenue(year int, month time.Month) int
}

func (or *OfficeReservation) getReservedDaysPerMonth(year int, month time.Month) int {
	return 0
}

func (or *OfficeReservation) getDailyPriceByMonth(year int, month time.Month) int {
	return 0
}

func (or *OfficeReservation) getMonthlyRevenue(year int, month time.Month) int {
	return 0
}
