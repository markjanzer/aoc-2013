package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"strconv"
)

const SmallTestString string = ``

const TestString string = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Can't move in one direction more than three times
	What is the way to get from the top left to the bottom right
	with the lowest heat loss (numbers) accumulated?

	Alright,
	We can calculate the score of a given set of moves by getting the average
	heat per positive direction. E and S are 1 direction, N and W are negative -1

	In order to be able to move and avoid squares with more heat loss, we need to be able to see more
	than one square ahead. For each move we'll look squaresAhead and determine the score for all of
	the possible outcomes. Then we'll take the best n (parallelTries) results, take the first
	move of them, and then execute those first moves


	To write this code, I'll start off with just checking one square ahead and doing no parallelization.


*/

type coordinates struct {
	x int
	y int
}

func oppositeDirection(direction string) string {
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
		panic("Invalid direction")
	}
}

func moveCoords(direction string, coords coordinates) coordinates {
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

func validDirections(grid [][]byte, currentCoords coordinates, path string) []string {
	directions := []string{"N", "E", "S", "W"}

	validDirections := []string{}

	for _, direction := range directions {
		if len(path) > 0 {
			// We cannot go the opposite direction of the last step
			lastDirection := string(path[len(path)-1])
			if direction == oppositeDirection(lastDirection) {
				continue
			}

			// We cannot go the same direction more than three times in a row
			if len(path) > 2 && direction == lastDirection && direction == string(path[len(path)-2]) && direction == string(path[len(path)-3]) {
				continue
			}
		}

		// We cannot go out of bounds of the grid
		newCoords := moveCoords(direction, currentCoords)
		if newCoords.x < 0 || newCoords.y < 0 || newCoords.x > len(grid[0])-1 || newCoords.y > len(grid)-1 {
			continue
		}

		validDirections = append(validDirections, direction)
	}

	return validDirections
}

type directionScores struct {
	direction string
	score     int
}

func heatLossAtCoords(grid [][]byte, coords coordinates) int {
	heatLoss, _ := strconv.Atoi(string(grid[coords.y][coords.x]))
	return heatLoss
}

func move(grid [][]byte, currentCoords coordinates, path string) (string, coordinates, int) {
	validDirections := validDirections(grid, currentCoords, path)

	bestScore := 100
	bestDirection := ""

	for _, direction := range validDirections {
		newCoords := moveCoords(direction, currentCoords)
		newHeatLoss := heatLossAtCoords(grid, newCoords)
		correctDirection := (newCoords.y - currentCoords.y) + (newCoords.x - currentCoords.x)
		score := newHeatLoss - (correctDirection * 5)

		if score < bestScore {
			bestScore = score
			bestDirection = direction
		}
	}

	newCoords := moveCoords(bestDirection, currentCoords)
	newHeatLoss := heatLossAtCoords(grid, newCoords)

	return bestDirection, newCoords, newHeatLoss
}

func printGridWithCoordinate(grid [][]byte, coord coordinates) {
	for y := range grid {
		for x := range grid[y] {
			if y == coord.y && x == coord.x {
				fmt.Print("X")
			} else {
				fmt.Print(string(grid[y][x]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func solvePart1(input string) int {
	grid := lib.StringToGrid(input)

	startingCoordinates := coordinates{0, 0}

	xMax := len(grid[0]) - 1
	yMax := len(grid) - 1

	currentCoords := startingCoordinates
	heatLost := 0
	path := ""
	for !(currentCoords.x == xMax && currentCoords.y == yMax) {
		newDirection, newCoords, newHeatLoss := move(grid, currentCoords, path)
		path += newDirection
		currentCoords = newCoords
		heatLost += newHeatLoss
		printGridWithCoordinate(grid, currentCoords)
	}

	return heatLost
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(102, solvePart1(TestString))
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
