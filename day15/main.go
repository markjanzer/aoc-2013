package main

import (
	"advent-of-code-2023/lib"
	"strconv"
	"strings"
)

const SmallTestString string = `HASH`
const SmallTestString2 string = `rn=1`

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

	Okay so first create a hash map that has numbers as the keys and arrays as the value
	Hmm I need to think about what the


*/

type lens struct {
	label       string
	focalLength int
}

type instruction struct {
	label       string
	focalLength int
	operation   string
}

// func (ins instruction) boxNumber() int {
// 	return hashString(ins.label)
// }

// var hashMap map[int][]lens

func parseInstruction(input string) instruction {
	if input[len(input)-1] == '-' {
		label := input[:len(input)-1]
		return instruction{label, 0, "-"}
	}

	strings := strings.Split(input, "=")
	label := strings[0]
	focalLength, _ := strconv.Atoi(strings[1])
	return instruction{label, focalLength, "="}
}

func solvePart2(input string) int {
	hashMap := make(map[int][]lens)

	// Populate the map with keys from 0 to 255
	for i := 0; i <= 255; i++ {
		hashMap[i] = []lens{}
	}

	inputs := strings.Split(input, ",")

	for _, value := range inputs {
		instruction := parseInstruction(value)
		boxNumber := hashString(instruction.label)

		if instruction.operation == "-" {
			indexOfLabel := lib.FindIndex(hashMap[boxNumber], func(lens lens) bool {
				return lens.label == instruction.label
			})
			if indexOfLabel != -1 {
				hashMap[boxNumber] = lib.RemoveIndex(hashMap[boxNumber], indexOfLabel)
			}
		} else if instruction.operation == "=" {
			indexOfLabel := lib.FindIndex(hashMap[boxNumber], func(lens lens) bool {
				return lens.label == instruction.label
			})
			if indexOfLabel != -1 {
				hashMap[boxNumber][indexOfLabel] = lens{instruction.label, instruction.focalLength}
			} else {
				hashMap[boxNumber] = append(hashMap[boxNumber], lens{instruction.label, instruction.focalLength})
			}
		}
	}

	sum := 0
	i := 0
	for i < 256 {
		for j, lens := range hashMap[i] {
			sum += (i + 1) * (j + 1) * lens.focalLength
		}
		i++
	}

	return sum
}

func main() {
	lib.AssertEqual(52, solvePart1(SmallTestString))
	lib.AssertEqual(1320, solvePart1(TestString))

	lib.AssertEqual(1, solvePart2(SmallTestString2))
	lib.AssertEqual(145, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
