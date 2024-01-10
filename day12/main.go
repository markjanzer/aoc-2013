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

// Takes a line and returns a map of the springs, and a log of the contiguous damaged springs
func getMapAndLog(input string) ([]byte, []int) {
	parts := strings.Split(input, " ")
	springsMap := []byte(parts[0])
	damagedSpringsLog := lib.Map(strings.Split(parts[1], ","), func(item string) int {
		num, _ := strconv.Atoi(item)
		return num
	})

	return springsMap, damagedSpringsLog
}

func solve(springsMap []byte, groups []int) int {
	// Break early if there are more damaged springs than the groups allow
	numberOfDamagedSprings := lib.Sum(groups...)
	damagedSprings := lib.Filter(springsMap, func(char byte) bool {
		return char == lib.CharToByte("#")
	})

	if len(damagedSprings) > numberOfDamagedSprings {
		return 0
	}

	// For each unknown spring, replace it with a damaged and undamaged spring and then run solve on those
	for i, char := range springsMap {
		if char == lib.CharToByte("?") {
			mapWithUndamagedSpring := make([]byte, len(springsMap))
			copy(mapWithUndamagedSpring, springsMap)
			mapWithUndamagedSpring[i] = lib.CharToByte(".")

			mapWithDamagedSpring := make([]byte, len(springsMap))
			copy(mapWithDamagedSpring, springsMap)
			mapWithDamagedSpring[i] = lib.CharToByte("#")

			return solve(mapWithUndamagedSpring, groups) + solve(mapWithDamagedSpring, groups)
		}
	}

	if springsMapMatchesPattern(springsMap, groups) {
		return 1
	}

	return 0
}

func generateSpringsLogForMap(springsMap []byte) (springsLog []int) {
	currentRange := 0
	for _, b := range springsMap {
		if string(b) == "#" {
			currentRange++
		} else {
			if currentRange != 0 {
				springsLog = append(springsLog, currentRange)
			}
			currentRange = 0
		}
	}

	if currentRange != 0 {
		springsLog = append(springsLog, currentRange)
	}

	return
}

func springsMapMatchesPattern(springsMap []byte, springsLog []int) bool {
	return lib.EqualSlices(generateSpringsLogForMap(springsMap), springsLog)
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
		row, groups := getMapAndLog(line)
		sum += solve(row, groups)
	}
	return
}

/*
	Part 2 Notes

	Make part 1 more efficient
	Okay there are two steps to this. Get a new getMapAndLog method that multiplies the results by 5

	How do I make part 1 more efficient?
	First off let's get some better naming to make this simpler. I like calling one part the map, and the other the
	log. Like logOfContiguousDamagedSprings.


*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(21, solvePart1(TestString))
	// lib.AssertEqual(525152, solvePart2(TestString))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
