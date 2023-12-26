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

const Part2TestString2 string = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

const Part2TestString3 string = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

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
	validDirections := []string{}
	for _, direction := range DIRECTIONS {
		if direction == tracker.CameFrom {
			continue
		} else {
			validDirections = append(validDirections, direction)
		}
	}

	for _, direction := range validDirections {
		if validMove(tracker.Coordinates, direction, tracker.Grid) {
			newTracker := Tracker{reverse(direction), tracker.Coordinates.move(direction), tracker.Distance + 1, tracker.Grid}
			return newTracker
		}
	}
	panic("No valid moves??")
}

func validMove(startingPoint Coordinates, direction string, grid [][]byte) bool {
	newCoordinates := startingPoint.move(direction)

	if !lib.PointInGrid(newCoordinates.X, newCoordinates.Y, grid) {
		return false
	}

	originalCharacter := string(grid[startingPoint.Y][startingPoint.X])
	newCharacter := string(grid[newCoordinates.Y][newCoordinates.X])

	return validMoveBetweenCharacters(direction, originalCharacter, newCharacter)
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

func validMoveBetweenCharacters(direction string, oldCharacter, newCharacter string) bool {
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

	tracker = tracker.move()
	for tracker.character() != "S" {
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


	Part 2 requires that we find the interior of the space created by the pipes

	What are some ways that we can do this?
	We could keep track of every single pipe in the route
	Then we could find the left topmost point and go to the right.
	In that row we have 0 space measured and a starting point, and loop := "open"
		In the next character over, if there is a space then we increment the total interior space and move on to the next item
		If the next character is a pipe then we say loop := "closed"

*/

func pipeUnderAnimal(coordinates Coordinates, grid [][]byte) string {
	validDirections := []string{}
	for _, direction := range DIRECTIONS {
		if validMove(coordinates, direction, grid) {
			validDirections = append(validDirections, direction)
		}
	}

	for character, directions := range characterFromMap {
		if character == "S" || character == "." {
			continue
		}
		// fmt.Println(validDirections, character, directions)
		if lib.All(directions, func(direction string) bool {
			return slices.Contains(validDirections, direction)
		}) {
			return character
		}
	}

	panic("No valid pipe found")
}

func solvePart2(input string) int {
	grid := lib.StringToGrid(input)
	gridCopy := grid
	var loopCoordinates = map[Coordinates]bool{}

	// fmt.Println("Previous Grid:")
	// printGrid(grid)

	animalCoordinates := findAnimal(grid)
	tracker := createTracker(animalCoordinates, grid)
	loopCoordinates[tracker.Coordinates] = true

	// tracker.print()

	tracker = tracker.move()
	loopCoordinates[tracker.Coordinates] = true

	for tracker.character() != "S" {
		tracker = tracker.move()
		loopCoordinates[tracker.Coordinates] = true
	}

	// Replace animal with pipe under animal
	gridCopy[animalCoordinates.Y][animalCoordinates.X] = lib.CharToByte(pipeUnderAnimal(animalCoordinates, grid))

	containedSpace := 0
	for y := range grid {
		loopOpen := false
		previousCorner := ""
		for x := range grid[y] {
			if loopCoordinates[Coordinates{X: x, Y: y}] {
				if gridCopy[y][x] == lib.CharToByte("|") {
					loopOpen = !loopOpen
				} else if gridCopy[y][x] == lib.CharToByte("-") {
					continue
				} else if isCorner(string(gridCopy[y][x])) {
					currentCorner := string(gridCopy[y][x])
					if previousCorner == "" {
						previousCorner = string(currentCorner)
					} else {
						if previousCorner == oppositeCorner(currentCorner) {
							loopOpen = !loopOpen
						}
						previousCorner = ""
					}
				}
			} else {
				if loopOpen {
					gridCopy[y][x] = lib.CharToByte("I")
					containedSpace += 1
				} else {
					gridCopy[y][x] = lib.CharToByte("0")
				}
			}
		}
	}

	// fmt.Println("New Grid:")
	// printGrid(gridCopy)
	// fmt.Println()

	return containedSpace
}

func oppositeCorner(pipe string) string {
	switch pipe {
	case "L":
		return "7"
	case "J":
		return "F"
	case "F":
		return "J"
	case "7":
		return "L"
	default:
		panic("Invalid pipe")
	}
}

func isCorner(pipe string) bool {
	return pipe == "L" || pipe == "J" || pipe == "F" || pipe == "7"
}

func printGrid(grid [][]byte) {
	for y := range grid {
		fmt.Println(string(grid[y]))
	}
}

func main() {
	// lib.AssertEqual(8, solvePart1(TestString))
	// lib.AssertEqual(4, solvePart1(SmallTestString))

	lib.AssertEqual(4, solvePart2(Part2TestString1))
	lib.AssertEqual(8, solvePart2(Part2TestString2))
	lib.AssertEqual(10, solvePart2(Part2TestString3))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// lib.AssertEqual(6842, result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
