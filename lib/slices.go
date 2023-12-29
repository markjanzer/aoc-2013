package lib

import "slices"

func LastValue[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Prepend[T any](collection []T, value T) []T {
	return append([]T{value}, collection...)
}

func IndexInSlice[T any](index int, slice []T) bool {
	return index >= 0 && index < len(slice)
}

func FindIndex[T any](collection []T, comparison func(T) bool) int {
	for i, value := range collection {
		if comparison(value) {
			return i
		}
	}
	return -1
}

func ContainsSameElements[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for _, value := range a {
		if !slices.Contains(b, value) {
			return false
		}
	}

	return true
}
