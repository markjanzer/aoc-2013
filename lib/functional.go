package lib

// Mostly stolen from lo
func Reduce[T any, R any](collection []T, reducer func(agg R, item T) R, initial R) R {
	for _, item := range collection {
		initial = reducer(initial, item)
	}
	return initial
}

func Map[T any, R any](collection []T, mapper func(item T) R) []R {
	result := []R{}
	for _, item := range collection {
		result = append(result, mapper(item))
	}
	return result
}

func FrequencyMap(input []string) map[string]int {
	frequencyMap := map[string]int{}
	for _, item := range input {
		frequencyMap[item]++
	}
	return frequencyMap
}

func Filter[T any](collection []T, filter func(item T) bool) []T {
	result := []T{}
	for _, item := range collection {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

func All[T any](collection []T, eval func(item T) bool) bool {
	for _, item := range collection {
		if !eval(item) {
			return false
		}
	}
	return true
}

func Any[T any](collection []T, eval func(item T) bool) bool {
	for _, item := range collection {
		if eval(item) {
			return true
		}
	}
	return false
}
