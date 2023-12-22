package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"slices"
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

const Part2TestString1 string = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const Part2TestString2 string = `...........`

const Part2TestString3 string = `...........`

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

// Determines whether you can move a direction from the character
var characterFromMap = map[string][]string{
	"S": DIRECTIONS,
	".": {},
	"|": {"N", "S"},
	"-": {"E", "W"},
	"L": {"N", "E"},
	"J": {"N", "W"},
	"F": {"S", "E"},
	"7": {"S", "W"},
}

// Determines whether you can move in a direction to the character
var characterToMap = map[string][]string{
	"S": DIRECTIONS,
	".": {},
	"|": {"N", "S"},
	"-": {"E", "W"},
	"L": {"S", "W"},
	"J": {"S", "E"},
	"F": {"N", "W"},
	"7": {"N", "E"},
}

func validFrom(direction, character string) bool {
	return slices.Contains(characterFromMap[character], direction)
}

func validTo(direction, character string) bool {
	return slices.Contains(characterToMap[character], direction)
}

func validMove(direction string, oldCharacter, newCharacter string) bool {
	return validTo(direction, newCharacter) && validFrom(direction, oldCharacter)
}

func reverse(direction string) string {
	indexOfDirection := lib.FindIndex(DIRECTIONS, func(value string) bool {
		return value == direction
	})
	return DIRECTIONS[(indexOfDirection+2)%4]
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
	lib.AssertEqual(4, solvePart1(SmallTestString))

	// lib.AssertEqual(4, solvePart2(Part2TestString1))
	// lib.AssertEqual(8, solvePart2(Part2TestString2))
	// lib.AssertEqual(10, solvePart2(Part2TestString3))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	lib.AssertEqual(6842, result1)
	fmt.Println(result1)

	// directions := []string{"N", "S"}
	// result := lib.Map(directions, func(direction string) string {
	// 	return reverse(direction)
	// })
	// fmt.Println(result)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
