package fractionrecurring

import (
	"slices"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {

	if numerator == 0 {
		return "0"
	}

	res := []byte{}
	if (numerator < 0) != (denominator < 0) {
		res = append(res, '-')
	}

	if numerator < 0 {
		numerator = -numerator
	}

	if denominator < 0 {
		denominator = -denominator
	}

	divRes := numerator / denominator
	res = append(res, []byte(strconv.Itoa(divRes))...)
	remainder := numerator % denominator

	if remainder == 0 {
		return string(res)
	}

	remainders := make(map[int]int)
	res = append(res, '.')

	for remainder != 0 {
		if index, exists := remainders[remainder]; exists {
			res = append(res, ')')
			res = slices.Insert(res, index, '(')
			break
		}

		remainders[remainder] = len(res)
		numerator = remainder * 10
		divRes = numerator / denominator
		remainder = numerator % denominator
		res = append(res, []byte(strconv.Itoa(divRes))...)
	}

	return string(res)
}
