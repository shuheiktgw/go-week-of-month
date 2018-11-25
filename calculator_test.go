package goWeekOfMonth

import (
	"testing"
	"time"
)

func TestGetWeekOfMonth(t *testing.T) {
	cases := []struct{
		year int
		month int
		day int
		want int
	} {
		{year: 2018, month: 11, day: 1, want: 1},
		{year: 2018, month: 11, day: 2, want: 1},
		{year: 2018, month: 11, day: 3, want: 1},
		{year: 2018, month: 11, day: 4, want: 1},
		{year: 2018, month: 11, day: 5, want: 2},
		{year: 2018, month: 11, day: 6, want: 2},
		{year: 2018, month: 11, day: 7, want: 2},
		{year: 2018, month: 11, day: 8, want: 2},
		{year: 2018, month: 11, day: 9, want: 2},
		{year: 2018, month: 11, day: 10, want: 2},
		{year: 2018, month: 11, day: 11, want: 2},
		{year: 2018, month: 11, day: 12, want: 3},
		{year: 2018, month: 11, day: 13, want: 3},
		{year: 2018, month: 11, day: 14, want: 3},
		{year: 2018, month: 11, day: 15, want: 3},
		{year: 2018, month: 11, day: 16, want: 3},
		{year: 2018, month: 11, day: 17, want: 3},
		{year: 2018, month: 11, day: 18, want: 3},
		{year: 2018, month: 11, day: 19, want: 4},
		{year: 2018, month: 11, day: 20, want: 4},
		{year: 2018, month: 11, day: 21, want: 4},
		{year: 2018, month: 11, day: 22, want: 4},
		{year: 2018, month: 11, day: 23, want: 4},
		{year: 2018, month: 11, day: 24, want: 4},
		{year: 2018, month: 11, day: 25, want: 4},
		{year: 2018, month: 11, day: 26, want: 5},
		{year: 2018, month: 11, day: 27, want: 5},
		{year: 2018, month: 11, day: 28, want: 5},
		{year: 2018, month: 11, day: 29, want: 5},
		{year: 2018, month: 11, day: 30, want: 5},
		{year: 2018, month: 12, day: 1, want: 1},
		{year: 2018, month: 12, day: 2, want: 1},
		{year: 2018, month: 12, day: 3, want: 2},
		{year: 2018, month: 12, day: 4, want: 2},
		{year: 2018, month: 12, day: 5, want: 2},
		{year: 2018, month: 12, day: 6, want: 2},
		{year: 2018, month: 12, day: 7, want: 2},
		{year: 2018, month: 12, day: 8, want: 2},
		{year: 2018, month: 12, day: 9, want: 2},
		{year: 2018, month: 12, day: 10, want: 3},
		{year: 2018, month: 12, day: 11, want: 3},
		{year: 2018, month: 12, day: 12, want: 3},
		{year: 2018, month: 12, day: 13, want: 3},
		{year: 2018, month: 12, day: 14, want: 3},
		{year: 2018, month: 12, day: 15, want: 3},
		{year: 2018, month: 12, day: 16, want: 3},
		{year: 2018, month: 12, day: 17, want: 4},
		{year: 2018, month: 12, day: 18, want: 4},
		{year: 2018, month: 12, day: 19, want: 4},
		{year: 2018, month: 12, day: 20, want: 4},
		{year: 2018, month: 12, day: 21, want: 4},
		{year: 2018, month: 12, day: 22, want: 4},
		{year: 2018, month: 12, day: 23, want: 4},
		{year: 2018, month: 12, day: 24, want: 5},
		{year: 2018, month: 12, day: 25, want: 5},
		{year: 2018, month: 12, day: 26, want: 5},
		{year: 2018, month: 12, day: 27, want: 5},
		{year: 2018, month: 12, day: 28, want: 5},
		{year: 2018, month: 12, day: 29, want: 5},
		{year: 2018, month: 12, day: 30, want: 5},
		{year: 2018, month: 12, day: 31, want: 1},
		{year: 2019, month: 1, day: 1, want: 1},
		{year: 2019, month: 1, day: 2, want: 1},
		{year: 2019, month: 1, day: 3, want: 1},
		{year: 2019, month: 1, day: 4, want: 1},
		{year: 2019, month: 1, day: 5, want: 1},
	}

	for i, tc := range cases {
		dt := time.Date(tc.year, time.Month(tc.month), tc.day, 1, 1, 1, 1, time.UTC)
		if got := GetWeekOfMonth(&dt); got != tc.want {
			t.Fatalf("#%d GetWeekOfMonth returns unexpected value: month: %v, day: %v, want: %v, got: %v", i, tc.month, tc.day, tc.want, got)
		}
	}
}

func TestGetWeekOfMonthByDay(t *testing.T) {
	cases := []struct{
		month int
		day int
		want int
	} {
		{month: 11, day: 1, want: 1},
		{month: 11, day: 2, want: 1},
		{month: 11, day: 3, want: 1},
		{month: 11, day: 4, want: 1},
		{month: 11, day: 5, want: 1},
		{month: 11, day: 6, want: 1},
		{month: 11, day: 7, want: 1},
		{month: 11, day: 8, want: 2},
		{month: 11, day: 9, want: 2},
		{month: 11, day: 10, want: 2},
		{month: 11, day: 11, want: 2},
		{month: 11, day: 12, want: 2},
		{month: 11, day: 13, want: 2},
		{month: 11, day: 14, want: 2},
		{month: 11, day: 15, want: 3},
		{month: 11, day: 16, want: 3},
		{month: 11, day: 17, want: 3},
		{month: 11, day: 18, want: 3},
		{month: 11, day: 19, want: 3},
		{month: 11, day: 20, want: 3},
		{month: 11, day: 21, want: 3},
		{month: 11, day: 22, want: 4},
		{month: 11, day: 23, want: 4},
		{month: 11, day: 24, want: 4},
		{month: 11, day: 25, want: 4},
		{month: 11, day: 26, want: 4},
		{month: 11, day: 27, want: 4},
		{month: 11, day: 28, want: 4},
		{month: 11, day: 29, want: 5},
		{month: 11, day: 30, want: 5},
		{month: 12, day: 1, want: 1},
		{month: 12, day: 2, want: 1},
		{month: 12, day: 3, want: 1},
		{month: 12, day: 4, want: 1},
		{month: 12, day: 5, want: 1},
		{month: 12, day: 6, want: 1},
		{month: 12, day: 7, want: 1},
		{month: 12, day: 8, want: 2},
		{month: 12, day: 9, want: 2},
		{month: 12, day: 10, want: 2},
		{month: 12, day: 11, want: 2},
		{month: 12, day: 12, want: 2},
		{month: 12, day: 13, want: 2},
		{month: 12, day: 14, want: 2},
		{month: 12, day: 15, want: 3},
		{month: 12, day: 16, want: 3},
		{month: 12, day: 17, want: 3},
		{month: 12, day: 18, want: 3},
		{month: 12, day: 19, want: 3},
		{month: 12, day: 20, want: 3},
		{month: 12, day: 21, want: 3},
		{month: 12, day: 22, want: 4},
		{month: 12, day: 23, want: 4},
		{month: 12, day: 24, want: 4},
		{month: 12, day: 25, want: 4},
		{month: 12, day: 26, want: 4},
		{month: 12, day: 27, want: 4},
		{month: 12, day: 28, want: 4},
		{month: 12, day: 29, want: 5},
		{month: 12, day: 30, want: 5},
		{month: 12, day: 31, want: 5},
	}

	for i, tc := range cases {
		dt := time.Date(2018, time.Month(tc.month), tc.day, 1, 1, 1, 1, time.UTC)
		if got := GetWeekOfMonthByDay(&dt); got != tc.want {
			t.Fatalf("#%d GetWeekOfMonthByDay returns unexpected value: month: %v, day: %v, want: %v, got: %v", i, tc.month, tc.day, tc.want, got)
		}
	}
}