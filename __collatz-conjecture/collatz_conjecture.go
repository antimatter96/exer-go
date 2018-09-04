package collatzconjecture

import (
	"errors"
)

// CollatzConjecture returns no of steps needed to get an integar to 1
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Number should be greater than equal to 1")
	}
	steps := 0
	for ; n != 1; steps++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}
	return steps, nil
}
