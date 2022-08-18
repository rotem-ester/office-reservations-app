package util

import (
	"testing"
	"time"
)

func TestGetDaysNumByYearAndMonth(t *testing.T) {
	type args struct {
		year  int
		month time.Month
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 29",
			args: args{
				year: 2020,
				month: time.February,
			},
			want: 29,
		},
		{
			name: "should return 31",
			args: args{
				year: 2020,
				month: time.December,
			},
			want: 31,
		},
		{
			name: "should return 31",
			args: args{
				year: 2020,
				month: time.January,
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDaysNumByYearAndMonth(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("GetDaysNumByYearAndMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
