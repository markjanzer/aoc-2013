package main

import (
	"advent-of-code-2023/lib"
	"fmt"
)

const SmallTestString string = ``

const TestString string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

const ShiftedTestString string = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Shift all of the Os north.
	Asign them points based on their vertical position on the grid
	and add up those points.

	How do I shift up?
	We can iterate over the grid space by space from top left to bottom right
	If we find a O we shift it up as many times as we can.
	Then we continue iterating over the grid.

	We could also go column by column. We could flip the grid.
	Then we would have arrays.
	Within these arrays we can iterate over the the spaces, and
	keep track of the last block (O or #).
*/

func solvePart1(input string) int {
	grid := lib.StringToGrid(input)
	grid = shiftNorth(grid)

	return countPoints(grid)
}

func shiftNorth(grid [][]byte) [][]byte {
	grid = lib.FlipGrid(grid)

	for y := range grid {
		block := -1
		for x := range grid[y] {
			if grid[y][x] == '#' {
				block = x
			} else if grid[y][x] == 'O' {
				if block < x {
					grid[y][x] = '.'
					grid[y][block+1] = 'O'
					block++
				} else {
					block = x
				}
			}
		}
	}

	grid = lib.FlipGrid(grid)

	return grid
}

func countPoints(grid [][]byte) int {
	points := 0
	gridHeight := len(grid)
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				points += gridHeight - y
			}
		}
	}

	return points
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(136, solvePart1(TestString))
	// lib.AssertEqual(136, solvePart1(TestString))
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
