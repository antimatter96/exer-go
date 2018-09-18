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

	return 0
}
