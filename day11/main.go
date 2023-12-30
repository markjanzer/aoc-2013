package main

import (
	"advent-of-code-2023/lib"
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
	return lib.MakeSlice(length, lib.CharToByte("."))
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
	return lib.AbsInt(a.X-b.X) + lib.AbsInt(a.Y-b.Y)
}

func expandedGalaxies(galaxies []Coordinates, gridWidth, gridHeight int, expansionFactor int) []Coordinates {
	spaceToAdd := expansionFactor - 1 // 1 * 100 = 1 + (100 - 1)
	xValues := lib.Map(galaxies, func(coordinates Coordinates) int {
		return coordinates.X
	})
	xGaps := lib.SliceDifference[int](lib.CreateRange(0, gridWidth-1), xValues)
	for i := len(xGaps) - 1; i >= 0; i-- {
		galaxies = lib.Map(galaxies, func(oldCoordinates Coordinates) Coordinates {
			if oldCoordinates.X > xGaps[i] {
				return Coordinates{oldCoordinates.X + spaceToAdd, oldCoordinates.Y}
			} else {
				return oldCoordinates
			}
		})
	}

	yValues := lib.Map(galaxies, func(coordinates Coordinates) int {
		return coordinates.Y
	})
	yGaps := lib.SliceDifference[int](lib.CreateRange(0, gridHeight-1), yValues)
	for i := len(yGaps) - 1; i >= 0; i-- {
		galaxies = lib.Map(galaxies, func(oldCoordinates Coordinates) Coordinates {
			if oldCoordinates.Y > yGaps[i] {
				return Coordinates{oldCoordinates.X, oldCoordinates.Y + spaceToAdd}
			} else {
				return oldCoordinates
			}
		})
	}

	return galaxies
}

func combinedGalaxyDistances(input string, expansionFactor int) int {
	grid := lib.StringToGrid(input)
	galaxies := galaxyCoordinates(grid)
	gridWidth := len(grid[0])
	gridHeight := len(grid)

	galaxies = expandedGalaxies(galaxies, gridWidth, gridHeight, expansionFactor)

	totalDistance := 0
	lib.CompareAllValues(galaxies, func(a, b Coordinates) {
		totalDistance += distanceBetween(a, b)
	})

	return totalDistance
}

func solvePart1(input string) int {
	return combinedGalaxyDistances(input, 2)
}

/*
	Part 2 Notes

	Alright, we're going to need a different strategy here.
	Instead of actually expanding the map here's what we're going to do.
	We're going to take the initial map and get the galaxies from it.
	Then we're going to iterate over the X values of the galaxies, and note any time there is a gap.
	Then we're going to increase the X values of all galaxies that are beyond that gap, as well as all of the gaps that are beyond it.
	Actually if we go over this backwards then we don't have to worry about expanding gaps.

	We're going to do the same thing for Y, then we're going to do the same distance calculation.

	For the first step, let's rewrite the first part to work this way.

*/

func solvePart2(input string) int {
	return combinedGalaxyDistances(input, 1_000_000)
}

func main() {
	// Expanding works!
	// lib.AssertEqual(TestStringExpanded, lib.GridToString(expandEmptyColumnsAndRows(lib.StringToGrid(TestString))))

	lib.AssertEqual(374, solvePart1(TestString))
	lib.AssertEqual(1030, combinedGalaxyDistances(TestString, 10))
	lib.AssertEqual(8410, combinedGalaxyDistances(TestString, 100))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	lib.AssertEqual(9742154, result1)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	result2 := solvePart2(dataString)
	lib.AssertEqual(411142919886, result2)
	// fmt.Println(result2)
}
