package main

import (
	"advent-of-code-2023/lib"
	"fmt"
)

const SmallTestString string = ``

const TestString string = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	I'm not quite sure the right way to to do this

	Okay so we can store the grid. Then we can have a laser that
	knows it's current coordinates.

	For determining whether a index has been energized we can use a
	map and serialize coordinates (comma separate), and have booleans
	Then we just get the number of keys and that's the result

	We do need laser.move() to be able to spawn more lasers though.

	I think we need different steps for the lasers. Actually it isn't
	even important that a laser exists continuously.

	We could have laser(grid, direction, coordinates)
*/

type coordinates struct {
	x int
	y int
}

func serialize(coordinates coordinates) string {
	return string(coordinates.x) + "," + string(coordinates.y)
}

func transformDirection(direction, tile string) []string {
	if tile == "." {
		return []string{direction}
	}

	if tile == "/" {
		switch direction {
		case "N":
			return []string{"E"}
		case "E":
			return []string{"N"}
		case "S":
			return []string{"W"}
		case "W":
			return []string{"S"}
		}
	}

	if tile == "\\" {
		switch direction {
		case "N":
			return []string{"W"}
		case "E":
			return []string{"S"}
		case "S":
			return []string{"E"}
		case "W":
			return []string{"N"}
		}
	}

	if tile == "|" {
		switch direction {
		case "N":
			return []string{"N"}
		case "E":
			return []string{"N", "S"}
		case "S":
			return []string{"S"}
		case "W":
			return []string{"N", "S"}
		}
	}

	if tile == "-" {
		switch direction {
		case "N":
			return []string{"E", "W"}
		case "E":
			return []string{"E"}
		case "S":
			return []string{"E", "W"}
		case "W":
			return []string{"W"}
		}
	}

	fmt.Println(tile)
	panic("Invalid tile")
}

func travel(direction string, coords coordinates) coordinates {
	switch direction {
	case "N":
		return coordinates{coords.x, coords.y - 1}
	case "E":
		return coordinates{coords.x + 1, coords.y}
	case "S":
		return coordinates{coords.x, coords.y + 1}
	case "W":
		return coordinates{coords.x - 1, coords.y}
	default:
		panic("Invalid direction")
	}
}

func coordinatesOutOfGrid(grid [][]byte, coordinates coordinates) bool {
	return coordinates.x < 0 || coordinates.y < 0 || coordinates.x >= len(grid[0]) || coordinates.y >= len(grid)
}

func laser(grid [][]byte, energizedTiles map[string]bool, direction string, coordinates coordinates) {
	if coordinatesOutOfGrid(grid, coordinates) {
		return
	}

	energizedTiles[serialize(coordinates)] = true

	tile := string(grid[coordinates.y][coordinates.x])
	newDirections := transformDirection(direction, tile)
	for _, newDirection := range newDirections {
		newCoordinates := travel(newDirection, coordinates)
		laser(grid, energizedTiles, newDirection, newCoordinates)
	}
}

func solvePart1(input string) int {
	grid := lib.StringToGrid(input)
	energizedTiles := make(map[string]bool)

	direction := "E"
	coordinates := coordinates{0, 0}

	laser(grid, energizedTiles, direction, coordinates)

	return len(energizedTiles)
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(46, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
