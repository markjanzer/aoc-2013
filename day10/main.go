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
	for _, direction := range DIRECTIONS {
		if direction != tracker.CameFrom && validMove(tracker.Coordinates, direction, tracker.Grid) {
			return tracker.moveDirection(direction)
		}
	}
	panic("No valid moves??")
}

func (tracker Tracker) moveDirection(direction string) Tracker {
	return Tracker{reverse(direction), tracker.Coordinates.move(direction), tracker.Distance + 1, tracker.Grid}
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

func characterAt(coordinates Coordinates, grid [][]byte) string {
	return string(grid[coordinates.Y][coordinates.X])
}

func (tracker Tracker) character() string {
	return characterAt(tracker.Coordinates, tracker.Grid)
}

func animalCoordinates(grid [][]byte) Coordinates {
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

	animalCoordinates := animalCoordinates(grid)
	tracker := createTracker(animalCoordinates, grid)

	for {
		tracker = tracker.move()
		if tracker.character() == "S" {
			break
		}
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

	We deterimine all of the points in the loop
	We figure out what the character is underneath the animal and replace the animal with it
	Then we iterate over each row in the grid
		We keep track of how many walls have been passed (walls are determined by N+S directions from a character)
		If we come across a non-loop character
			If there are an odd number of walls in front of the character
				It is inside and we count it
			Otherwise it is outside and we don't count it

	Return the sum
*/

func pipeUnderAnimal(coordinates Coordinates, grid [][]byte) string {
	fromDirectionsOfPipe := []string{}
	for _, direction := range DIRECTIONS {
		if validMove(coordinates, direction, grid) {
			fromDirectionsOfPipe = append(fromDirectionsOfPipe, direction)
		}
	}

	for character, directions := range characterFromMap {
		if character != "S" && character != "." && lib.ContainsSameElements(fromDirectionsOfPipe, directions) {
			return character
		}
	}

	panic("No valid pipe found")
}

func replaceAnimal(coordinates Coordinates, grid [][]byte) {
	grid[coordinates.Y][coordinates.X] = lib.CharToByte(pipeUnderAnimal(coordinates, grid))
}

func resetPreviousWalls() map[string]int {
	return map[string]int{
		"N": 0,
		"S": 0,
		"E": 0,
		"W": 0,
	}
}

// func printGrid(grid [][]byte) {
// 	for y := range grid {
// 		fmt.Println(string(grid[y]))
// 	}
// }

func solvePart2(input string) int {
	grid := lib.StringToGrid(input)
	var isPartOfLoop = map[Coordinates]bool{}

	// fmt.Println("Previous Grid:")
	// printGrid(grid)

	animalCoordinates := animalCoordinates(grid)
	tracker := createTracker(animalCoordinates, grid)

	for {
		isPartOfLoop[tracker.Coordinates] = true
		tracker = tracker.move()
		if tracker.character() == "S" {
			break
		}
	}

	replaceAnimal(animalCoordinates, grid)

	containedSpace := 0
	for y := range grid {
		wallsInFrontOfPoint := 0
		previousWalls := resetPreviousWalls()
		for x := range grid[y] {
			if isPartOfLoop[Coordinates{X: x, Y: y}] {
				character := string(grid[y][x])
				for _, direction := range characterFromMap[character] {
					previousWalls[direction]++
				}
				if previousWalls["N"] == 1 && previousWalls["S"] == 1 {
					wallsInFrontOfPoint++
					previousWalls = resetPreviousWalls()
				}
				if previousWalls["S"] >= 2 || previousWalls["N"] >= 2 {
					previousWalls = resetPreviousWalls()
				}
			} else {
				if wallsInFrontOfPoint%2 == 1 {
					containedSpace++
					grid[y][x] = lib.CharToByte("I")
				} else {
					grid[y][x] = lib.CharToByte("0")
				}
			}
		}
	}

	return containedSpace
}

func main() {
	lib.AssertEqual(8, solvePart1(TestString))
	lib.AssertEqual(4, solvePart1(SmallTestString))

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
