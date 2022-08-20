package office_reservation

import (
	"time"

	util "github.com/rotem-ester/office-reservation-app/service/pkg/util"
)

const DATE_LAYOUT = "2006-01-02"

type (
	OfficeReservation struct {
		Capacity     int
		MonthlyPrice int
		StartDay     time.Time
		EndDay       time.Time
	}
)

func (or *OfficeReservation) getMonthlyRevenue(year int, month time.Month) int {
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
	return float64(or.MonthlyPrice) / days
}
