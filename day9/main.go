package main

import (
	"advent-of-code-2023/lib"
	"strings"
)

const SmallTestString string = `10 13 16 21 30 45`

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

func createPredictionStructure(input string) (predictionStructure [][]int) {
	firstArray := lib.IntsFromString(input)
	predictionStructure = append(predictionStructure, firstArray)

	for {
		lastSlice := lib.LastValue(predictionStructure)
		newSlice := differences(lastSlice)
		predictionStructure = append(predictionStructure, newSlice)

		allValuesAreZero := lib.All(newSlice, func(value int) bool {
			return value == 0
		})

		if allValuesAreZero {
			break
		}
	}

	return predictionStructure
}

func differences(numbers []int) []int {
	result := []int{}
	for i := 1; i < len(numbers); i++ {
		difference := numbers[i] - numbers[i-1]
		result = append(result, difference)
	}
	return result
}

func expandForward(predictionStructure [][]int) {
	expandedValue := 0
	for i := len(predictionStructure) - 1; i >= 0; i-- {
		currentRow := predictionStructure[i]
		expandedValue += lib.LastValue(currentRow)
		predictionStructure[i] = append(currentRow, expandedValue)
	}
}

func lastTopLevelValue(predictionStructure [][]int) int {
	return predictionStructure[0][len(predictionStructure[0])-1]
}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		predictionStructure := createPredictionStructure(line)
		expandForward(predictionStructure)
		value := lastTopLevelValue(predictionStructure)
		sum += value
	}

	return sum
}

/*
	Part 2 Notes

	Instead of extrapolating the future, we're extrapolating the past
	A similar solution should work for the first values and subtraction rather than last values and addition
*/

func expandBackward(predictionStructure [][]int) {
	expandedValue := 0
	for i := len(predictionStructure) - 1; i >= 0; i-- {
		currentRow := predictionStructure[i]
		expandedValue = currentRow[0] - expandedValue
		predictionStructure[i] = lib.Prepend(currentRow, expandedValue)
	}
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		predictionStructure := createPredictionStructure(line)
		expandBackward(predictionStructure)
		value := predictionStructure[0][0]
		sum += value
	}

	return sum
}

func main() {
	lib.AssertEqual(114, solvePart1(TestString))
	lib.AssertEqual(2, solvePart2(TestString))

	lib.AssertEqual(5, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
