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
