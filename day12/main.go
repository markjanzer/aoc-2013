package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strconv"
	"strings"
)

const SmallTest1 string = `???.### 1,1,3`           // => 1
const SmallTest2 string = `.??..??...?##. 1,1,3`    // => 4
const SmallTest3 string = `?#?#?#?#?#?#?#? 1,3,1,6` // => 1

const TestString string = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Okay first we get the row and the groups (of damaged springs), and we pass that to
	Then we iterate through the string with a solve method
	First this method checks if the given string has more damaged springs than provided in
		the groups, and if it does then it returns 0
	This method iterates over the string. When it comes across a ?
		It replaces the character with a . and with a #, and it returns the sum of solve being called
		with those two strings
	If it is not a ? then continue.
	It returns 1

*/

func getRowAndGroups(input string) ([]byte, []int) {
	parts := strings.Split(input, " ")
	row := []byte(parts[0])
	groups := lib.Map(strings.Split(parts[1], ","), func(item string) int {
		num, _ := strconv.Atoi(item)
		return num
	})

	fmt.Println("groups: ", groups)

	return row, groups
}

func solve(row []byte, groups []int) int {
	printByteSlice(row)

	// Break early if there are more damaged springs than the groups allow
	expectedDamagedSprings := lib.Sum(groups...)
	damagedSprings := lib.Filter(row, func(char byte) bool {
		return char == lib.CharToByte("#")
	})

	if len(damagedSprings) > expectedDamagedSprings {
		return 0
	}

	// For each unknown spring, replace it with a damaged and undamaged spring and then run solve on those
	for i, char := range row {
		if char == lib.CharToByte("?") {
			undamagedSpring := make([]byte, len(row))
			copy(undamagedSpring, row)
			undamagedSpring[i] = lib.CharToByte(".")

			damagedSpring := make([]byte, len(row))
			copy(damagedSpring, row)
			damagedSpring[i] = lib.CharToByte("#")

			return solve(undamagedSpring, groups) + solve(damagedSpring, groups)
		}
	}

	if rowMatchesPattern(row, groups) {
		return 1
	}

	return 0
}

func generatePatternForRow(row []byte) (pattern []int) {
	currentRange := 0
	for _, b := range row {
		if string(b) == "#" {
			currentRange++
		} else {
			if currentRange != 0 {
				pattern = append(pattern, currentRange)
			}
			currentRange = 0
		}
	}

	if currentRange != 0 {
		pattern = append(pattern, currentRange)
	}

	return
}

func rowMatchesPattern(row []byte, pattern []int) bool {
	return lib.EqualSlices(generatePatternForRow(row), pattern)
}

func printByteSlice(b []byte) {
	var stringsSlice []string
	for _, byteVal := range b {
		stringsSlice = append(stringsSlice, string(byteVal))
	}
	fmt.Println(stringsSlice)
}

func solvePart1(input string) (sum int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row, groups := getRowAndGroups(line)
		sum += solve(row, groups)
	}
	return
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(21, solvePart1(TestString))
	// lib.AssertEqual(21, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
