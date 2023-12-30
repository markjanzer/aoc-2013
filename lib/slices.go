package lib

import (
	"slices"

	"golang.org/x/exp/constraints"
)

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

func CompareAllValues[T any](collection []T, compare func(a, b T)) {
	for i := 0; i < len(collection)-1; i++ {
		for j := i + 1; j < len(collection); j++ {
			compare(collection[i], collection[j])
		}
	}
}

func MakeSlice[T any](length int, value T) []T {
	slice := make([]T, length)
	for i := range slice {
		slice[i] = value
	}
	return slice
}

// Returns the elements that are in slice A but not in slice B
// SliceDifference([]int{1,2,3,4,5}, []int{1,3,4}) => []int{2,5}
func SliceDifference[T constraints.Ordered](a, b []T) []T {
	slices.Sort(a)
	slices.Sort(b)

	var difference []T
	aIndex := 0
	bIndex := 0
	for aIndex < len(a) && bIndex < len(b) {
		if a[aIndex] == b[bIndex] {
			aIndex++
			bIndex++
		} else if a[aIndex] < b[bIndex] {
			difference = append(difference, a[aIndex])
			aIndex++
		} else if a[aIndex] > b[bIndex] {
			bIndex++
		}
	}

	difference = append(difference, a[aIndex:]...)

	return difference
}

func CreateRange(start, end int) []int {
	slice := make([]int, 0, end-start)
	for i := start; i <= end; i++ {
		slice = append(slice, i)
	}
	return slice
}
