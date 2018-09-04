// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
