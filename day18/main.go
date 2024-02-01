package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strconv"
	"strings"
)

const SmallTestString string = ``

const TestString string = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Ah first we have to get grid dimensions which can be a little tricky.
	We also will need the coordinates of all of the points for the next step!

	So how about this. We assume that the start is at 0,0. We'll create a
	coordinate for that point, and each time we'll create a coordinate we'll update the
	colMin, colMax, rowMin, and rowMax coordinates.
	Then we can determine the dimensions of the grid with colMax-colMin and rowMax-rowMin,
	And we can shift all of the coordinates by the -colMin and -rowMin

	Then we can create a grid and iterate over it



	Ah also a little tricky is creating the map from the given instructions.
	We could assume that the start is in the top right, but that might not be correct.

	We ignore the hexcodes in this part. What we do here is we map out the path,
	and then we iterate over the grid, counting the edges, and interior.
	We sum edges and interior squares to get the result.

*/

type instruction struct {
	direction string
	distance  int
}

func getInstructions(input string) (instructions []instruction) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, instruction{direction, distance})
	}

	return
}

type coordinates struct {
	row int
	col int
}

func move(coords coordinates, direction string) coordinates {
	switch direction {
	case "U":
		return coordinates{coords.row - 1, coords.col}
	case "D":
		return coordinates{coords.row + 1, coords.col}
	case "L":
		return coordinates{coords.row, coords.col - 1}
	case "R":
		return coordinates{coords.row, coords.col + 1}
	default:
		panic("Invalid direction")
	}
}

const Edge = "#"
const Space = "."

func createGridFromInstructions(instructions []instruction) [][]byte {
	minCol := 0
	maxCol := 0
	minRow := 0
	maxRow := 0

	currentCoords := coordinates{0, 0}
	edges := []coordinates{currentCoords}

	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			currentCoords = move(currentCoords, instruction.direction)
			// Don't add the coordinates again when reaching the beginning
			if currentCoords.col == 0 && currentCoords.row == 0 {
				continue
			}
			edges = append(edges, currentCoords)
		}

		if currentCoords.col < minCol {
			minCol = currentCoords.col
		} else if currentCoords.col > maxCol {
			maxCol = currentCoords.col
		}
		if currentCoords.row < minRow {
			minRow = currentCoords.row
		} else if currentCoords.row > maxRow {
			maxRow = currentCoords.row
		}
	}

	edges = lib.Map(edges, func(coord coordinates) coordinates {
		return coordinates{coord.row - minRow, coord.col - minCol}
	})

	newRowMax := maxRow - minRow
	newColMax := maxCol - minCol

	grid := lib.CreateGrid(newRowMax+1, newColMax+1, lib.CharToByte(Space))

	for _, edge := range edges {
		grid[edge.row][edge.col] = lib.CharToByte(Edge)
	}

	return grid
}

func sumEdgesAndInterior(grid [][]byte) int {
	gridToPrint := lib.CreateGrid(len(grid), len(grid[0]), lib.CharToByte(Space))

	edgeSquares := 0
	interiorSquares := 0
	for row := range grid {
		edgesPassed := 0
		previousSquareWasEdge := false
		for col := range grid[row] {
			switch string(grid[row][col]) {
			case Edge:
				if !previousSquareWasEdge {
					edgesPassed++
				}
				edgeSquares++
				gridToPrint[row][col] = lib.CharToByte(Edge)
				previousSquareWasEdge = true
			case Space:
				if edgesPassed%2 != 0 {
					interiorSquares++
					gridToPrint[row][col] = lib.CharToByte(Edge)
				}
				previousSquareWasEdge = false
			default:
				fmt.Println("Invalid character in grid")
				fmt.Println(grid[row][col])
			}
		}
	}

	lib.PrintGrid(gridToPrint)

	return edgeSquares + interiorSquares
}

func solvePart1(input string) int {
	instructions := getInstructions(input)
	grid := createGridFromInstructions(instructions)

	lib.PrintGrid(grid)
	fmt.Println()
	return sumEdgesAndInterior(grid)
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	// lib.AssertEqual(62, solvePart1(TestString))
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
