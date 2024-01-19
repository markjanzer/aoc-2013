package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strings"
)

const SmallTestString string = `HASH`

const TestString string = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Write hash algorithm that takes a string and returns an number
	Separate the string by commas and then sum all of the hashed values

	To write the hashing algorithm
*/

func hashString(input string) int {
	currentValue := 0
	for _, char := range input {
		currentValue = hashChar(string(char), currentValue)
	}
	return currentValue
}

func hashChar(char string, currentValue int) int {
	currentValue += int(char[0])
	currentValue *= 17
	currentValue %= 256

	return currentValue
}

func solvePart1(input string) int {
	inputs := strings.Split(input, ",")

	result := 0
	for _, value := range inputs {
		result += hashString(value)
	}

	return result
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(52, solvePart1(SmallTestString))
	lib.AssertEqual(1320, solvePart1(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
