package main

import (
	"advent-of-code-2023/lib"
	"fmt"
)

const SmallTestString string = `.....
.S-7.
.|.|.
.L-J.
.....`

const TestString string = `
..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	We're going to start by making a grid of bytes
	Then we're going to find the S, and get the current coordinates

	Okay when travelling we need to know the character we're coming from, the direction we're going
	and the character we're going to

*/

type Tracker struct {
	CameFrom    string
	Coordinates Coordinates
	Distance    int
	Grid        [][]byte
}

type Coordinates struct {
	X int
	Y int
}

func (coordinates Coordinates) move(direction string) Coordinates {
	newX, newY := coordinates.X, coordinates.Y
	switch direction {
	case "N":
		newY -= 1
	case "E":
		newX += 1
	case "S":
		newY += 1
	case "W":
		newX -= 1
	default:
		panic("we done goofed")
	}
	return Coordinates{newX, newY}
}

var DIRECTIONS = []string{"N", "E", "S", "W"}

func (tracker Tracker) move() Tracker {
	for _, direction := range DIRECTIONS {
		if direction == tracker.CameFrom {
			continue
		}

		fmt.Println()
		fmt.Println("direction", direction)
		newCoordinates := tracker.Coordinates.move(direction)

		fmt.Println("newCoordinates", newCoordinates)

		if !lib.PointInGrid(newCoordinates.X, newCoordinates.Y, tracker.Grid) {
			continue
		}

		fmt.Println("valid point in grid")

		newCharacter := tracker.characterAt(newCoordinates)

		fmt.Println(direction, newCharacter)

		if validMove(direction, tracker.character(), newCharacter) {
			fmt.Println("valid move!")
			newTracker := Tracker{reverse(direction), newCoordinates, tracker.Distance + 1, tracker.Grid}
			return newTracker
		}
	}
	panic("No valid moves??")
}

// Determines whether or not a move a certain direction can
// be made to the character

func validTo(direction, character string) bool {
	switch character {
	case "S":
		return true
	case ".":
		return false
	case "|":
		return direction == "N" || direction == "S"
	case "-":
		return direction == "E" || direction == "W"
	case "L":
		return direction == "S" || direction == "W"
	case "J":
		return direction == "S" || direction == "E"
	case "7":
		return direction == "N" || direction == "E"
	case "F":
		return direction == "N" || direction == "W"
	default:
		fmt.Println(character)
		panic("Not a valid character")
	}
}

func validFrom(direction, character string) bool {
	switch character {
	case "S":
		return true
	case ".":
		panic("Moving from .")
	case "|":
		return direction == "N" || direction == "S"
	case "-":
		return direction == "E" || direction == "W"
	case "L":
		return direction == "N" || direction == "E"
	case "J":
		return direction == "N" || direction == "W"
	case "7":
		return direction == "S" || direction == "W"
	case "F":
		return direction == "S" || direction == "E"
	default:
		fmt.Println(character)
		panic("Not a valid character")
	}
}

func validMove(direction string, oldCharacter, newCharacter string) bool {
	return validTo(direction, newCharacter) && validFrom(direction, oldCharacter)
}

func reverse(direction string) string {
	switch direction {
	case "N":
		return "S"
	case "E":
		return "W"
	case "S":
		return "N"
	case "W":
		return "E"
	default:
		fmt.Println(direction)
		panic("Not a valid direction")
	}
}

func (tracker Tracker) characterAt(coordinates Coordinates) string {
	return string(tracker.Grid[coordinates.Y][coordinates.X])
}

func (tracker Tracker) character() string {
	return tracker.characterAt(tracker.Coordinates)
}

func findAnimal(grid [][]byte) Coordinates {
	for y := range grid {
		for x := range grid[y] {
			if string(grid[y][x]) == "S" {
				return Coordinates{x, y}
			}
		}
	}
	panic("Animal not found!")
}

func createTracker(coordinates Coordinates, grid [][]byte) Tracker {
	return Tracker{"", coordinates, 0, grid}
}

func solvePart1(input string) int {
	grid := lib.StringToGrid(input)

	animalCoordinates := findAnimal(grid)
	tracker := createTracker(animalCoordinates, grid)

	tracker.print()

	tracker = tracker.move()
	for tracker.character() != "S" {
		tracker.print()
		tracker = tracker.move()
	}

	return tracker.Distance / 2
}

func (tracker Tracker) print() {
	fmt.Println()
	fmt.Println("Tracker")
	fmt.Println("Current Character", tracker.character())
	fmt.Println("Distance: ", tracker.Distance)
	fmt.Println("X: ", tracker.Coordinates.X)
	fmt.Println("Y: ", tracker.Coordinates.Y)
	fmt.Println("CameFrom: ", tracker.CameFrom)
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(8, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	lib.AssertEqual(4, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
