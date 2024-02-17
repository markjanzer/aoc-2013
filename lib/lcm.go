package lib

import (
	"golang.org/x/exp/constraints"
)

func Lcm[T constraints.Integer](a, b T) T {
	return a * b / Gcd(a, b)
}

func Gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LcmOfSlice[T constraints.Integer](numbers []T) T {
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = Lcm(result, num)
	}
	return result
}
