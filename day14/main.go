package main

import (
	"advent-of-code-2023/lib"
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
	grid = shiftNorth(grid)

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

	It does seem like it permanently settles, but rather oscilates between a couple of states

	We'll start by serializing the grids
	Then we'll save the results of cycles in a map (with the serialized value as the key and the cycle number as the value)
	When we find a duplicate result, we'll calculate the difference to determine if there is a consistent
	cycle it reappears with and see if 1_000_000_000 appears in that cycle. If it does, then calculate the result for the grid
*/

func cycleGrid(grid [][]byte) [][]byte {
	grid = shiftNorth(grid)
	grid = rotateGrid(grid)
	grid = shiftNorth(grid)
	grid = rotateGrid(grid)
	grid = shiftNorth(grid)
	grid = rotateGrid(grid)
	grid = shiftNorth(grid)
	grid = rotateGrid(grid)
	return grid
}

// Rotates grid 90 degrees counter-clockwise
func rotateGrid(grid [][]byte) [][]byte {
	xLen := len(grid[0])
	yLen := len(grid)

	newGrid := make([][]byte, xLen)
	for i := range newGrid {
		newGrid[i] = make([]byte, yLen)
	}

	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			newGrid[x][yLen-1-y] = grid[y][x]
		}
	}

	return newGrid
}

func serialize(grid [][]byte) string {
	return lib.GridToString(grid)
}

func solvePart2(input string) int {
	gridMap := make(map[string]int)
	grid := lib.StringToGrid(input)

	cycle := 0
	for cycle < 10000 {
		grid = cycleGrid(grid)
		cycle++

		serializedGrid := serialize(grid)

		if _, ok := gridMap[serializedGrid]; ok {
			previousCycle := gridMap[serializedGrid]
			period := cycle - previousCycle
			if (1_000_000_000-previousCycle)%period == 0 {
				return countPoints(grid)
			}
		}

		gridMap[serializedGrid] = cycle
	}

	panic("No solution found")
}

func main() {
	lib.AssertEqual(136, solvePart1(TestString))
	lib.AssertEqual(64, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
