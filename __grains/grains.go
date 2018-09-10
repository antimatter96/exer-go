package grains

import "errors"

// Square returns the no of grains at the nth square (1 based indexing)
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n should be between 1 and 64")
	}
	var power = uint(n)
	return 1 << (power - 1), nil
}

// Total returns the no of grains on the whole board (with 64 squares)
func Total() uint64 {
	return (1 << 64) - 1
}
