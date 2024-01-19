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

const NorthShiftedTestString string = `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`

const CycledTestString string = `.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`

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
	lib.PrintGrid(grid)
	fmt.Println()
	grid = shiftNorth(grid)
	lib.PrintGrid(grid)

	return countPoints(grid)
}

func shiftNorth(grid [][]byte) [][]byte {
	for x := range grid[0] {
		block := -1
		for y := range grid {
			if grid[y][x] == '#' {
				block = y
			} else if grid[y][x] == 'O' {
				if block < y {
					grid[y][x] = '.'
					grid[block+1][x] = 'O'
					block++
				} else {
					block = y
				}
			}
		}
	}

	return grid
}

// func shiftNorth(grid [][]byte) [][]byte {
// 	grid = lib.FlipGrid(grid)
// 	grid = shiftWest(grid)
// 	grid = lib.FlipGrid(grid)

// 	return grid
// }

// func shiftWest(grid [][]byte) [][]byte {
// 	for y := range grid {
// 		block := -1
// 		for x := range grid[y] {
// 			if grid[y][x] == '#' {
// 				block = x
// 			} else if grid[y][x] == 'O' {
// 				if block < x {
// 					grid[y][x] = '.'
// 					grid[y][block+1] = 'O'
// 					block++
// 				} else {
// 					block = x
// 				}
// 			}
// 		}
// 	}

// 	return grid
// }

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

	We need to do a cycle of shifting North, West, South, then East
	We also need a method to compare grids

	We'll do cycles and compare it to the previous result until it settles.

	Then we'll once again calculate points.
*/

// func cycleGrid(grid [][]byte) [][]byte {
// 	// North
// 	grid = lib.FlipGrid(grid)
// 	grid = shiftWest(grid)
// 	// West
// 	grid = lib.FlipGrid(grid)
// 	grid = shiftWest(grid)
// 	// South
// 	grid = lib.FlipGrid(grid)
// 	grid = lib.ReverseGrid(grid)
// 	grid = shiftWest(grid)
// 	// East
// 	grid = lib.FlipGrid(grid)
// 	grid = shiftWest(grid)
// 	// Restore grid
// 	grid = lib.ReverseGrid(grid)

// 	return grid
// }

func cycleGrid(grid [][]byte) [][]byte {
	grid = shiftNorth(grid)
	grid = rotateGrid(grid, counterClockwise)
	grid = shiftNorth(grid)
	grid = rotateGrid(grid, counterClockwise)
	grid = shiftNorth(grid)
	grid = rotateGrid(grid, counterClockwise)
	grid = shiftNorth(grid)
	grid = rotateGrid(grid, counterClockwise)
	return grid
}

// func shiftSouth(grid [][]byte) [][]byte {
// 	grid = lib.FlipGrid(grid)
// 	grid = lib.ReverseGrid(grid)
// 	grid = shiftWest(grid)
// 	grid = lib.ReverseGrid(grid)
// 	grid = lib.FlipGrid(grid)
// 	return grid
// }

// func shiftEast(grid [][]byte) [][]byte {
// 	grid = lib.ReverseGrid(grid)
// 	grid = shiftWest(grid)
// 	grid = lib.ReverseGrid(grid)
// 	return grid
// }

// Rotates grid 90 degrees counter-clockwise
func rotateGrid(grid [][]byte, direction string) [][]byte {
	xLen := len(grid[0])
	yLen := len(grid)

	newGrid := make([][]byte, xLen)
	for i := range newGrid {
		newGrid[i] = make([]byte, yLen)
	}

	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if direction == "cc" {
				newGrid[x][yLen-1-y] = grid[y][x]
			} else if direction == "c" {
				newGrid[xLen-1-x][y] = grid[y][x]
			}
		}
	}

	return newGrid
}

func solvePart2(input string) int {
	grid := lib.StringToGrid(input)
	lib.PrintGrid(grid)
	fmt.Println()

	return 0
}

const clockwise = "c"
const counterClockwise = "cc"

func main() {
	// lib.AssertEqual(136, solvePart1(TestString))
	// lib.AssertEqual(136, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// 	numbers := `123
	// 456
	// 789`
	// 	numGrid := lib.StringToGrid(numbers)
	// 	lib.PrintGrid(numGrid)
	// 	fmt.Println()
	// 	numGrid = rotateGrid(numGrid, clockwise)
	// 	lib.PrintGrid(numGrid)
	// 	fmt.Println()
	// 	numGrid = rotateGrid(numGrid, clockwise)
	// 	numGrid = rotateGrid(numGrid, counterClockwise)
	// 	lib.PrintGrid(numGrid)

	grid := lib.StringToGrid(TestString)
	// cycledGrid := cycleGrid(grid)
	// expectedGrid := lib.StringToGrid(CycledTestString)

	i := 0
	for i < 10000 {
		grid = cycleGrid(grid)
		i++
	}
	lib.PrintGrid(grid)
	fmt.Println()

	i = 0
	for i < 1000 {
		grid = cycleGrid(grid)
		i++
	}
	lib.PrintGrid(grid)

	// fmt.Println(lib.GridAreEqual(cycledGrid, expectedGrid))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
