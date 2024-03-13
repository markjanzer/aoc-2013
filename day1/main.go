package main

import (
	"advent-of-code-2023/lib"
	"strings"
	"unicode"
)

const FirstTestString string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const SecondTestString string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

const DataFile string = "data.txt"

func solvePart1(input string) int {
	calibrationCodes := strings.Split(input, "\n")

	sum := 0
	for _, line := range calibrationCodes {
		digits := firstAndLastDigit1(line)
		sum += ((digits[0] * 10) + digits[1])
	}

	return sum
}

func firstAndLastDigit1(line string) []int {
	var digits = make([]int, 0, 2)

	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			digits = appendDigit(digits, intFromUnicode(runeValue))
		}
	}

	return digits
}

func appendDigit(digits []int, digit int) []int {
	if len(digits) == 0 {
		digits = append(digits, digit)
		digits = append(digits, digit)
	} else {
		digits[1] = digit
	}
	return digits
}

func intFromUnicode(runeValue rune) int {
	return int(runeValue - '0')
}

// Part 2

var digitNames = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func solvePart2(input string) int {
	calibrationCodes := strings.Split(input, "\n")

	sum := 0
	for _, line := range calibrationCodes {
		digits := firstAndLastDigit2(line)
		sum += ((digits[0] * 10) + digits[1])
	}

	return sum
}

func firstAndLastDigit2(line string) []int {
	var digits = make([]int, 0, 2)
	var sequentialLetters string = ""

	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			sequentialLetters = ""
			digits = appendDigit(digits, intFromUnicode(runeValue))
		} else {
			sequentialLetters += string(runeValue)
			var count = len(sequentialLetters)
			for name, value := range digitNames {
				if count >= len(name) && sequentialLetters[count-len(name):] == name {
					digits = appendDigit(digits, value)
				}
			}
		}
	}

	return digits
}

func main() {
	lib.AssertEqual(142, solvePart1(FirstTestString))
	lib.AssertEqual(281, solvePart2(SecondTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
