package clock

import (
	"fmt"
)

// A Clock value contains hours and minutes
type Clock struct {
	minutes int
	hours   int
}

// New returns a Clock value with the hour, minute set
func New(hour, minute int) Clock {
	for minute < 0 {
		hour--
		minute += 60
	}
	if minute >= 60 {
		hour += minute / 60
		minute %= 60
	}
	for hour < 0 {
		hour += 24
	}
	hour = (hour % 24)
	return Clock{minute, hour}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

// Add returns a clock object with the given minutes added
func (c Clock) Add(minutes int) Clock {
	return New(c.hours, c.minutes+minutes)
}

// Subtract returns a clock object with the given minutes subtracted
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hours, c.minutes-minutes)
}
