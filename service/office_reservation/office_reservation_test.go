package office_reservation

import (
	"testing"
	"time"
)

func TestOfficeReservation_getReservedDaysByMonth(t *testing.T) {
	type fields struct {
		Capacity     int
		MonthlyPrice int
		StartDay     time.Time
		EndDay       time.Time
	}
	type args struct {
		year  int
		month time.Month
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "June 2014 - should return 30",
			fields: fields{
				Capacity:     14,
				MonthlyPrice: 11875,
				StartDay:     time.Date(2014, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.June,
			},
			want: 30,
		},
		{
			name: "July 2014 - should return 31",
			fields: fields{
				Capacity:     14,
				MonthlyPrice: 11875,
				StartDay:     time.Date(2014, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.July,
			},
			want: 31,
		},
		{
			name: "May 2014 - should return 0",
			fields: fields{
				Capacity:     14,
				MonthlyPrice: 11875,
				StartDay:     time.Date(2014, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.May,
			},
			want: 0,
		},
		{
			name: "April 2014 - should return 29",
			fields: fields{
				Capacity:     1,
				MonthlyPrice: 400,
				StartDay:     time.Date(2014, time.April, 2, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.April,
			},
			want: 29,
		},
		{
			name: "April 2014 - should return 22",
			fields: fields{
				Capacity:     3,
				MonthlyPrice: 1850,
				StartDay:     time.Date(2014, time.April, 9, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Date(2014, time.August, 6, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				year:  2014,
				month: time.April,
			},
			want: 22,
		},
		{
			name: "May 2014 - should return 31",
			fields: fields{
				Capacity:     3,
				MonthlyPrice: 1850,
				StartDay:     time.Date(2014, time.April, 9, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Date(2014, time.August, 6, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				year:  2014,
				month: time.May,
			},
			want: 31,
		},
		{
			name: "August 2014 - should return 6",
			fields: fields{
				Capacity:     3,
				MonthlyPrice: 1850,
				StartDay:     time.Date(2014, time.April, 9, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Date(2014, time.August, 6, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				year:  2014,
				month: time.August,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or := &OfficeReservation{
				Capacity:     tt.fields.Capacity,
				MonthlyPrice: tt.fields.MonthlyPrice,
				StartDay:     tt.fields.StartDay,
				EndDay:       tt.fields.EndDay,
			}
			if got := or.getReservedDaysByMonth(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("OfficeReservation.getReservedDaysByMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfficeReservation_GetMonthlyRevenue(t *testing.T) {
	type fields struct {
		Capacity     int
		MonthlyPrice int
		StartDay     time.Time
		EndDay       time.Time
	}
	type args struct {
		year  int
		month time.Month
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "June 2014 - should return 11875",
			fields: fields{
				Capacity:     14,
				MonthlyPrice: 11875,
				StartDay:     time.Date(2014, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.June,
			},
			want: 11875,
		},
		{
			name: "July 2014 - should return 11875",
			fields: fields{
				Capacity:     14,
				MonthlyPrice: 11875,
				StartDay:     time.Date(2014, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.July,
			},
			want: 11875,
		},
		{
			name: "May 2014 - should return 0",
			fields: fields{
				Capacity:     14,
				MonthlyPrice: 11875,
				StartDay:     time.Date(2014, time.June, 1, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.May,
			},
			want: 0,
		},
		{
			name: "April 2014 - should return 386",
			fields: fields{
				Capacity:     1,
				MonthlyPrice: 400,
				StartDay:     time.Date(2014, time.April, 2, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Time{},
			},
			args: args{
				year:  2014,
				month: time.April,
			},
			want: 386,
		},
		{
			name: "April 2014 - should return 1356",
			fields: fields{
				Capacity:     3,
				MonthlyPrice: 1850,
				StartDay:     time.Date(2014, time.April, 9, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Date(2014, time.August, 6, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				year:  2014,
				month: time.April,
			},
			want: 1356,
		},
		{
			name: "May 2014 - should return 1850",
			fields: fields{
				Capacity:     3,
				MonthlyPrice: 1850,
				StartDay:     time.Date(2014, time.April, 9, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Date(2014, time.August, 6, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				year:  2014,
				month: time.May,
			},
			want: 1850,
		},
		{
			name: "August 2014 - should return 358",
			fields: fields{
				Capacity:     3,
				MonthlyPrice: 1850,
				StartDay:     time.Date(2014, time.April, 9, 0, 0, 0, 0, time.UTC),
				EndDay:       time.Date(2014, time.August, 6, 0, 0, 0, 0, time.UTC),
			},
			args: args{
				year:  2014,
				month: time.August,
			},
			want: 358,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or := &OfficeReservation{
				Capacity:     tt.fields.Capacity,
				MonthlyPrice: tt.fields.MonthlyPrice,
				StartDay:     tt.fields.StartDay,
				EndDay:       tt.fields.EndDay,
			}
			if got := or.GetMonthlyRevenue(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("OfficeReservation.GetMonthlyRevenue() = %v, want %v", got, tt.want)
			}
		})
	}
}
