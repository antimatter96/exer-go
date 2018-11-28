package twelve

import (
	"fmt"
	"strings"
)

var items = []string{
	"Partridge in a Pear Tree",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

var countOfItems = []string{
	"a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
}
var days = []string{
	"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

// generateSecondPart creates the part after "gave to me: "
func generateSecondPart(n int) string {
	var stringBuilder strings.Builder
	for index := n - 1; index > 0; index-- {
		stringBuilder.WriteString(fmt.Sprintf("%v %v, ", countOfItems[index], items[index]))
	}
	toRet := stringBuilder.String()
	if len(toRet) > 0 {
		toRet = toRet + "and "
	}
	toRet = toRet + countOfItems[0] + " " + items[0] + "."
	return toRet
}

// Verse return the line at day input
func Verse(input int) string {
	if input < 1 || input > 12 {
		return "Error not a valid input, must be from 1 to 12"
	}

	firstPart := fmt.Sprintf("On the %v day of Christmas my true love gave to me: ", days[input-1])
	secondPart := generateSecondPart(input)
	return firstPart + secondPart
}

// Song returns the whole song. Internally calls Verse
func Song() string {
	var stringBuilder strings.Builder
	for index := 1; index < 12+1; index++ {
		stringBuilder.WriteString(Verse(index))
		stringBuilder.WriteByte(10)
	}
	return stringBuilder.String()
}
