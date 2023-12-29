package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"slices"
)

const TestString string = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

const TestStringExpanded string = `....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Turn the string into a grid
	Expand it by adding an empty row after finding an existing empty row
	Flip the grid, then do it again
	Then flip it back (for clarity + error checking purposes)

	Find the coordinates for every galaxy

	Create a method to find the distance between two galaxies

	Compare all of the elements, and add the distnaces in a sum
*/

const EmptyChar = "."
const GalaxyChar = "#"

func expandEmptyRows(grid [][]byte) [][]byte {
	y := 0
	for y < len(grid) {
		if rowIsEmpty(grid[y]) {
			grid = slices.Insert(grid, y, emptyRow(len(grid[y])))
			y++
		}
		y++
	}
	return grid
}

func rowIsEmpty(row []byte) bool {
	return lib.All(row, func(item byte) bool {
		return item == lib.CharToByte(EmptyChar)
	})
}

func emptyRow(length int) []byte {
	return makeSlice(length, lib.CharToByte("."))
}

func makeSlice(length int, value byte) []byte {
	slice := make([]byte, length)
	for i := range slice {
		slice[i] = value
	}
	return slice
}

func expandEmptyColumnsAndRows(grid [][]byte) [][]byte {
	grid = lib.FlipGrid(grid)
	grid = expandEmptyRows(grid)

	// Expand rows
	grid = lib.FlipGrid(grid)
	grid = expandEmptyRows(grid)

	return grid
}

type Coordinates struct {
	X int
	Y int
}

func galaxyCoordinates(grid [][]byte) (coordinates []Coordinates) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == lib.CharToByte(GalaxyChar) {
				newCoordinate := Coordinates{x, y}
				coordinates = append(coordinates, newCoordinate)
			}
		}
	}
	return
}

func distanceBetween(a, b Coordinates) int {
	return absInt(a.X-b.X) + absInt(a.Y-b.Y)
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func compareAllValues(collection []Coordinates, compare func(a, b Coordinates)) {
	for i := 0; i < len(collection)-1; i++ {
		for j := i + 1; j < len(collection); j++ {
			compare(collection[i], collection[j])
		}
	}
}

func solvePart1(input string) int {
	grid := lib.StringToGrid(input)
	grid = expandEmptyColumnsAndRows(grid)
	galaxies := galaxyCoordinates(grid)

	totalDistance := 0
	compareAllValues(galaxies, func(a, b Coordinates) {
		totalDistance += distanceBetween(a, b)
	})

	return totalDistance
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	// Expanding works!
	// lib.AssertEqual(TestStringExpanded, lib.GridToString(expandEmptyColumnsAndRows(lib.StringToGrid(TestString))))

	lib.AssertEqual(374, solvePart1(TestString))

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
