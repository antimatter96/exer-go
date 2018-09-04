// Package gigasecond is used to provide operations dealing with time and Gigaseconds
package gigasecond

// import path for the time package from the standard library
import "time"

var oneGigaSecond = time.Second * 1000000000

// AddGigasecond adds One Gigaseconds to the provided time and returns it
func AddGigasecond(t time.Time) time.Time {
	return t.Add(oneGigaSecond)
}
