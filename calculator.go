package goWeekOfMonth

import "time"

// GetWeekOfMonth returns the number of the week of month. Weeks start on Monday.
// For example, 2018/11/05 belongs to the second week of November so GetWeekOfMonth returns two
func GetWeekOfMonth(t *time.Time) int {
	startOfMonth := time.Date(t.Year(), t.Month(), 1, 1, 1, 1, 1, t.Location())
	_, currentWeek := t.ISOWeek()
	_, firstWeek := startOfMonth.ISOWeek()

	w := 1 + currentWeek - firstWeek

	// Dec 29 to Dec 31 might belong to week 1 of year n+1
	if w < 0 {
		return 1
	}

	return w
}

// GetWeekOfMonthByDay returns the number of the week of month by day
// For example, 2018/11/05 is the first Monday in 2018/11, so GetWeekOfMonthByDay returns one
func GetWeekOfMonthByDay(t *time.Time) int {
	startOfMonth := time.Date(t.Year(), t.Month(), 1, 1, 1, 1, 1, t.Location())
	days := t.Sub(startOfMonth).Hours() / 24
	return 1 + int(days / 7)
}