package perfect

import (
	"errors"
	"math"
)

type Classification int

var ErrOnlyPositive error = errors.New("Only positive numbers")

const (
	ClassificationPerfect   = 1
	ClassificationAbundant  = 2
	ClassificationDeficient = 3
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil
	}

	var sum int64

	sqrt := int64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		sum += sqrt
	} else {
		sqrt++
	}

	sum++
	for i := int64(2); i < sqrt; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}

	if sum == n {
		return ClassificationPerfect, nil
	}

	if sum > n {
		return ClassificationAbundant, nil
	}

	return ClassificationDeficient, nil
}
