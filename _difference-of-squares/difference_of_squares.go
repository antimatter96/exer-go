package diffsquares

// SquareOfSums returns the square of the sum of the first n natural numbers
func SquareOfSums(n int) int {
	sum := (n * (n + 1) / 2)
	return sum * sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * ((2 * n) + 1) / 6)
}

/*
Difference returns the difference in
the square of the sum of the first n natural numbers and
the sum of the squares of the first n natural numbers
*/
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
