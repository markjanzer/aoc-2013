package lib

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
