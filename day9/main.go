package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strings"
)

const SmallTestString string = ``

const TestString string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	For each row get the next extrapolated value
	Add them all together

	We're going to store these values in an array as the first value of a 2D array
	If all of the values are not 0, add another array that diffs it.

	Once we create an array with all of the values as 0 we extrapolate:
	Append last array with 0
	Append the previous array with it's last value + this arrays last value
	When there is no higher array, return that value
*/

func extrapolatedValue(input string) int {
	structure := [][]int{}
	firstArray := lib.IntsFromString(input)
	structure = append(structure, firstArray)

	for {
		lastSlice := lib.LastValue(structure)
		newSlice := differences(lastSlice)
		structure = append(structure, newSlice)

		allValuesAreZero := lib.All(newSlice, func(value int) bool {
			return value == 0
		})

		if allValuesAreZero {
			break
		}
	}

	return sumOfLastValues(structure)
}

func differences(numbers []int) []int {
	result := []int{}
	for i := 1; i < len(numbers); i++ {
		difference := numbers[i] - numbers[i-1]
		result = append(result, difference)
	}
	return result
}

func sumOfLastValues(structure [][]int) int {
	sum := 0
	for _, slice := range structure {
		sum += lib.LastValue(slice)
	}
	return sum
}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")

	return lib.Reduce(lines, func(result int, line string) int {
		return result + extrapolatedValue(line)
	}, 0)
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(114, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
