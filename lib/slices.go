package lib

func LastValue[T any](slice []T) T {
	return slice[len(slice)-1]
}
