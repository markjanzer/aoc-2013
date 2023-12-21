package lib

func LastValue[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Prepend[T any](collection []T, value T) []T {
	return append([]T{value}, collection...)
}
