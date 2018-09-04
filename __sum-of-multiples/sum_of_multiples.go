package summultiples

// SumMultiples returns the sum of numbers less than `limit` divisible by `divisors`
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, j := range divisors {
			if i%j == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
