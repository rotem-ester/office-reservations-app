package main

import "time"

type (
	OfficeReservation struct {
		Capacity int
		MonthlyPrice int
		StartDay time.Time
		EndDay time.Time
	}
)