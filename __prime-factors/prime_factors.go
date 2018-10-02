package prime

func Factors(n int64) []int64 {
	primeFactors := make([]int64, 0)

	var i int64 = 2

	for n%i == 0 {
		primeFactors = append(primeFactors, i)
		n /= i
	}

	i = 3
	for n > 1 {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n /= i
		}
		i += 2
	}

	return primeFactors
}
