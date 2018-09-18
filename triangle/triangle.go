package triangle

import "math"

type Kind int

const (
	NaT = -1 // not a triangle
	Equ = 3  // equilateral
	Iso = 2  // isosceles
	Sca = 1  // scalene
)

// KindFromSides return the kind of triangle it is given all three sides
func KindFromSides(a, b, c float64) Kind {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if a == b && b == c && a == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	return Sca
}
