package allyourbase

import (
	"errors"
	"math/big"
)

// ConvertToBase converts the number in form of array of digits in given input base to given output base
func ConvertToBase(inputBase int, digits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	inpBase := big.NewInt(int64(inputBase))
	outBase := big.NewInt(int64(outputBase))
	currentPower := big.NewInt(1)
	decimalRep := big.NewInt(0)
	toAdd := big.NewInt(0)

	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] < 0 || digits[i] >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		} else if digits[i] > 0 {
			toAdd = big.NewInt(int64(digits[i]))
			toAdd = toAdd.Mul(toAdd, currentPower)
			decimalRep = decimalRep.Add(decimalRep, toAdd)
		}
		currentPower = currentPower.Mul(currentPower, inpBase)
	}

	toRet := make([]int, 0)

	m := big.NewInt(1)
	for decimalRep.Int64() != 0 {
		decimalRep, m = decimalRep.DivMod(decimalRep, outBase, m)
		toRet = append(toRet, int(m.Int64()))
	}

	for i, j := 0, len(toRet)-1; i < j; i, j = i+1, j-1 {
		toRet[i], toRet[j] = toRet[j], toRet[i]
	}

	if len(toRet) == 0 {
		toRet = append(toRet, 0)
	}

	return toRet, nil
}
