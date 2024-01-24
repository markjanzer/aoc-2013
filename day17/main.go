package main

import (
	"advent-of-code-2023/lib"
	"fmt"
	"sort"
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


	Ah the scoring is difficult

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

func (state travelState) validDirections() []string {
	// No valid directions if you are at the end
	if state.distanceFromEnd() == 0 {
		return []string{}
	}

	directions := []string{"N", "E", "S", "W"}
	validDirections := []string{}

	for _, direction := range directions {
		if len(state.path) > 0 {
			// We cannot go the opposite direction of the last step
			lastDirection := string(state.path[len(state.path)-1])
			if direction == oppositeDirection(lastDirection) {
				continue
			}
			// We cannot go the same direction more than three times in a row
			if len(state.path) > 2 && direction == lastDirection && direction == string(state.path[len(state.path)-2]) && direction == string(state.path[len(state.path)-3]) {
				continue
			}
		}
		// We cannot go out of bounds of the grid
		newCoords := moveCoords(direction, state.coordinates)
		if newCoords.x < 0 || newCoords.y < 0 || newCoords.x > len(state.grid[0])-1 || newCoords.y > len(state.grid)-1 {
			continue
		}
		validDirections = append(validDirections, direction)
	}

	return validDirections
}

func (state travelState) distanceFromEnd() int {
	xMax := len(state.grid[0]) - 1
	yMax := len(state.grid) - 1

	return (xMax - state.coordinates.x) + (yMax - state.coordinates.y)
}

func (state travelState) completed() bool {
	return state.distanceFromEnd() == 0
}

func (state travelState) score() float64 {
	// maxDistance := len(state.grid) + len(state.grid[0])
	// distanceTravelled := maxDistance - state.distanceFromEnd()
	return float64(state.heatLoss) + (float64(state.distanceFromEnd()) * 5)
}

func heatLossAtCoords(grid [][]byte, coords coordinates) int {
	heatLoss, _ := strconv.Atoi(string(grid[coords.y][coords.x]))
	return heatLoss
}

type travelState struct {
	grid        [][]byte
	coordinates coordinates
	path        string
	heatLoss    int
}

func move(direction string, state travelState) travelState {
	newCoords := moveCoords(direction, state.coordinates)
	newHeatLoss := state.heatLoss + heatLossAtCoords(state.grid, newCoords)
	newPath := state.path + direction

	return travelState{state.grid, newCoords, newPath, newHeatLoss}
}

func moveAllValidDirections(state travelState, squaresAhead int, resultingStates *[]travelState) {
	if squaresAhead == 0 || state.distanceFromEnd() == 0 {
		*resultingStates = append(*resultingStates, state)
		return
	}

	for _, direction := range state.validDirections() {
		newState := move(direction, state)
		moveAllValidDirections(newState, squaresAhead-1, resultingStates)
	}
}

func (state travelState) printGrid() {
	coords := []coordinates{{0, 0}}
	for i, direction := range state.path {
		coords = append(coords, moveCoords(string(direction), coords[i]))
	}

	for y := range state.grid {
		for x := range state.grid[y] {
			if lib.Any(coords, func(coord coordinates) bool { return coord.x == x && coord.y == y }) {
				fmt.Print("X")
			} else {
				fmt.Print(string(state.grid[y][x]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

const squaresAhead = 10
const parallelTries = 100

func solvePart1(input string) int {
	grid := lib.StringToGrid(input)

	states := []travelState{{grid, coordinates{0, 0}, "", 0}}
	bestCompletedState := travelState{grid, coordinates{0, 0}, "", 0}

	counter := 0
	for counter < 1000 {
		resultingStates := []travelState{}
		for _, state := range states {
			moveAllValidDirections(state, squaresAhead, &resultingStates)
		}

		for _, state := range resultingStates {
			if state.completed() && (bestCompletedState.score() == 0 || state.score() < bestCompletedState.score()) {
				bestCompletedState = state
			}
		}

		resultingStates = lib.Filter(resultingStates, func(state travelState) bool { return !state.completed() })

		sort.Slice(resultingStates, func(i, j int) bool {
			return resultingStates[i].score() < resultingStates[j].score()
		})

		if bestCompletedState.score() != 0 && bestCompletedState.score() < resultingStates[0].score() {
			break
		}

		states = resultingStates[:parallelTries]

		// fmt.Println("Coords", state.coordinates, "Path", state.path, "Heat Loss", state.heatLoss)
		// state.printGrid()

		counter++
		fmt.Println(counter)
	}

	return bestCompletedState.heatLoss
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

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
