package main

import (
	"advent-of-code-2023/lib"
)

const SmallTestString string = `.....
.S-7.
.|.|.
.L-J.
.....`

const TestString string = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const DataFile string = "data.txt"

/*
	Part 1 Notes

*/

func solvePart1(input string) int {
	return 0
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(4, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	lib.AssertEqual(8, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
