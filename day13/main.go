package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strings"
)

const SmallTestString string = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`

const TestString string = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Formatting:
	Iterate over the string reading each matrix as a grid
	Get solution for each grid

	Solution:
	For each grid if there is horizontal symmetry, return the number of
	columns before symmetry occurs.
	If there is vertical symmetry, return the number of rows before the
	symmetry occurs, and then multiply by 100

	How to determine horizontal symmetry:
	Compare each row to the last. If effective then go one row back and forward
	for each and compare again. Do this until one of the indexes is at the max or 0
	If it fails then keep iterating

	To check for vertical symmetry, I'll flip the grid and then check for horizontal symmetry
*/

func solvePart1(input string) (sum int) {
	inputs := strings.Split(input, "\n\n")

	for _, input := range inputs {
		sum += solvePart1ForGrid(input)
	}

	return
}

func solvePart1ForGrid(input string) int {
	grid := lib.StringToGrid(input)
	// return horizontalSymmetryAfterRow(grid)
	return (100 * horizontalSymmetryAfterRow(grid)) + verticalSymmetryAfterColumn(grid)
}

func verticalSymmetryAfterColumn(grid [][]byte) int {
	flippedGrid := lib.FlipGrid(grid)
	return horizontalSymmetryAfterRow(flippedGrid)
}

func horizontalSymmetryAfterRow(grid [][]byte) int {
	for i := 1; i < len(grid); i++ {
		// Continue if symmetry is not found
		prev := i - 1
		if !lib.EqualSlices(grid[i], grid[prev]) {
			continue
		}

		// If symmetry is found, continue looking forward and
		// back to compare rows until not possible
		symmetrical := true
		symmetryDistance := 1
		for (i+symmetryDistance) < len(grid) && (prev-symmetryDistance) >= 0 {
			if !lib.EqualSlices(grid[i+symmetryDistance], grid[prev-symmetryDistance]) {
				symmetrical = false
				break
			}
			symmetryDistance++
		}

		if symmetrical {
			return i
		}
	}

	return 0
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(5, solvePart1(SmallTestString))
	lib.AssertEqual(405, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart2(SmallTestString))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
