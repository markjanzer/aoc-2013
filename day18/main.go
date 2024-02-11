package main

import (
	"advent-of-code-2023/lib"
	"strconv"
	"strings"
)

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

	Using the shoelace formula
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

func moveDistance(coords coordinates, direction string, distance int) coordinates {
	switch direction {
	case "U":
		return coordinates{coords.row - distance, coords.col}
	case "D":
		return coordinates{coords.row + distance, coords.col}
	case "L":
		return coordinates{coords.row, coords.col - distance}
	case "R":
		return coordinates{coords.row, coords.col + distance}
	default:
		panic("Invalid direction")
	}
}

const Edge = "#"
const Space = "."

func createCoordsFromInstructions(instructions []instruction) ([]coordinates, int) {
	currentCoords := coordinates{0, 0}
	coords := []coordinates{currentCoords}
	borderSize := 0

	for _, instruction := range instructions {
		currentCoords = moveDistance(currentCoords, instruction.direction, instruction.distance)
		coords = append(coords, currentCoords)
		borderSize += instruction.distance
	}

	return coords, borderSize
}

func shoelaceFormula(coords []coordinates) int {
	sum := 0
	for i := 0; i < len(coords)-1; i++ {
		currentCoords := coords[i]
		nextCoords := coords[i+1]
		sum += currentCoords.row*nextCoords.col - currentCoords.col*nextCoords.row
	}

	return lib.AbsInt(sum) / 2
}

// Calculate the number of interior squares using Pick's theorem
func interiorSquares(area, borderSize int) int {
	return area - (borderSize / 2) + 1
}

func solvePart1(input string) int {
	instructions := getInstructions(input)
	coordinates, borderSize := createCoordsFromInstructions(instructions)
	area := shoelaceFormula(coordinates)
	interiorSquares := interiorSquares(area, borderSize)
	return interiorSquares + borderSize
}

/*
	Part 2 Notes

	Turn hexadecimal into numbers
	Run the same solution (but probably needs to be more efficient)
*/

var digitToDirection = map[string]string{
	"0": "R",
	"1": "D",
	"2": "L",
	"3": "U",
}

func getInstructionsPart2(input string) (instructions []instruction) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, "#")
		// Get the the hex without the trailing parenthesis
		fullHex := parts[1][:len(parts[1])-1]
		numberHex := fullHex[:len(fullHex)-1]
		direction := digitToDirection[string(fullHex[len(fullHex)-1])]
		distance, _ := strconv.ParseInt(numberHex, 16, 64)
		intDistance := int(distance)

		instructions = append(instructions, instruction{direction, intDistance})
	}

	return
}

func solvePart2(input string) int {
	instructions := getInstructionsPart2(input)
	coordinates, borderSize := createCoordsFromInstructions(instructions)
	area := shoelaceFormula(coordinates)
	interiorSquares := interiorSquares(area, borderSize)
	return interiorSquares + borderSize
}

func main() {
	lib.AssertEqual(62, solvePart1(TestString))
	lib.AssertEqual(952408144115, solvePart2(TestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
}
