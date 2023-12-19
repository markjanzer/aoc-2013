package lib

import (
	"strconv"
	"strings"
)

// Takes a string like " 79 14 55 13 " and returns a slice of ints like [79, 14, 55, 13]
func IntsFromString(input string) (result []int) {
	fields := strings.Fields(input)
	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		result = append(result, num)
	}
	return
}

// Returns whether or not integer is within given range
func IntIsInRange(input, begin, end int) bool {
	return input >= begin && input <= end
}
