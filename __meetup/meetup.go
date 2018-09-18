package meetup

import "time"

// WeekSchedule is
type WeekSchedule int

//
var (
	First  WeekSchedule = 1
	Second WeekSchedule = 2
	Third  WeekSchedule = 3
	Fourth WeekSchedule = 5
	Last   WeekSchedule = -1
	Teenth WeekSchedule = 10
)

// Day i
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	if week == Last {
		return last(weekday, month, year)
	}
	if week == Teenth {
		return teenth(weekday, month, year)
	}
	firstDate := first(weekday, month, year)
	if week == First {
		return firstDate
	}
	if week == Second {
		return firstDate + 7
	}
	if week == Third {
		return firstDate + 14
	}
	if week == Fourth {
		return firstDate + 21
	}
	return 0
}

func last(weekday time.Weekday, month time.Month, year int) int {
	for d := 31; d > 0; d-- {
		t := time.Date(year, month, d, 23, 0, 0, 0, time.UTC)
		if t.Weekday() == weekday {
			if month == time.February {
				if d > 29 {
					continue
				}
				if d == 29 {
					if !isLeapYear(year) {
						continue
					}
				}
			}
			return d
		}
	}
	return 0
}

func teenth(weekday time.Weekday, month time.Month, year int) int {
	for d := 13; d < 20; d++ {
		t := time.Date(year, month, d, 23, 0, 0, 0, time.UTC)
		if t.Weekday() == weekday {

			return d
		}
	}
	return 0
}

func first(weekday time.Weekday, month time.Month, year int) int {
	for d := 1; d < 31; d++ {
		t := time.Date(year, month, d, 23, 0, 0, 0, time.UTC)
		if t.Weekday() == weekday {
			return d
		}
	}
	return 0
}

func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
