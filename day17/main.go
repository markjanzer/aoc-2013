package main

import (
	"advent-of-code-2023/lib"
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
	Part 1 attempt 2

	Alright I need two things to make this work, a heap and then an A* algorithm that uses that heap
	I made the heap. Now I'm going to make this work.

	Let's try to describe what we want to build here

	We're going to take the original state
	We're going to find all of the valid directions. Then we're going to move to that square.
	When we move to that square we'll store a few things in the new state
	- The coordinates they are on
	- The cost of getting there
	- The move they took to get there
	- The number of repeated moves
	We're going to have a map with the squares as keys, and the cost to get there as the value
	Whenever we move to a square, we see if the map has a cost value that is lower than the current cost
		If it is lower than the current cost then we don't do anything, we've already found a more
		efficient way to get to this square
	If the current cost is lower or there is no lowest value for that square then we set that value in the map
	Then we determine the priority for this state. We'll have an int where lower is better, and it will be current
	cost + distance from the end
	Then we add this square to the heap (which determines order by the priority)

	I guess that this starts with a state in the top left with everything calculated, and we put that in the heap.
	Then we pop from the heap, get a list of valid directions, then travel to each of those, checking the
	new square against the map, and then potentially adding it to the heap.


*/

type coordinates struct {
	x int
	y int
}

type travelState struct {
	coords        coordinates
	difficulty    int
	lastDirection string
	wentDirection int
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
		// This is to handle the initial state, not sure if I like it here
		return "none"
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

func (state travelState) validDirections(grid [][]byte) []string {
	// No valid directions if you are at the end
	if distanceFromEnd(grid, state.coords) == 0 {
		return []string{}
	}

	directions := []string{"N", "E", "S", "W"}
	validDirections := []string{}

	for _, direction := range directions {
		// We cannot go the opposite direction of the last step
		if direction == oppositeDirection(state.lastDirection) {
			continue
		}
		// We cannot go the same direction more than three times in a row
		if state.wentDirection > 3 {
			continue
		}
		// We cannot go out of bounds of the grid
		newCoords := moveCoords(direction, state.coords)
		if newCoords.x < 0 || newCoords.y < 0 || newCoords.x > len(grid[0])-1 || newCoords.y > len(grid)-1 {
			continue
		}
		validDirections = append(validDirections, direction)
	}

	return validDirections
}

func distanceFromEnd(grid [][]byte, coords coordinates) int {
	xMax := len(grid[0]) - 1
	yMax := len(grid) - 1
	return (xMax - coords.x) + (yMax - coords.y)
}

func serializeCoords(coords coordinates) string {
	return strconv.Itoa(coords.x) + "," + strconv.Itoa(coords.y)
}

func addCoordsDifficulty(coordsDifficulty map[string]int, state travelState) bool {
	key := serializeCoords(state.coords)
	if originalDifficulty, exists := coordsDifficulty[key]; !exists || originalDifficulty > state.difficulty {
		coordsDifficulty[key] = state.difficulty
		return true
	}
	return false
}

func difficultyAt(grid [][]byte, coords coordinates) int {
	return lib.IntFromByte(grid[coords.y][coords.x])
}

func (state travelState) move(grid [][]byte, direction string) travelState {
	newCoords := moveCoords(direction, state.coords)
	newDifficulty := state.difficulty + difficultyAt(grid, newCoords)
	var newWentDirection int
	if direction == state.lastDirection {
		newWentDirection = state.wentDirection + 1
	} else {
		newWentDirection = 1
	}
	return travelState{newCoords, newDifficulty, direction, newWentDirection}
}

func solvePart1(input string) int {
	grid := lib.StringToGrid(input)

	upNext := lib.NewHeap(func(a, b travelState) bool {
		return distanceFromEnd(grid, a.coords) < distanceFromEnd(grid, b.coords)
	})

	coordsDifficulty := make(map[string]int)
	initialState := travelState{coordinates{0, 0}, 0, "none", 0}

	addCoordsDifficulty(coordsDifficulty, initialState)
	upNext.Insert(initialState)

	for upNext.Size() > 0 {
		state := upNext.Pop()

		for _, direction := range state.validDirections(grid) {
			newState := state.move(grid, direction)
			if addCoordsDifficulty(coordsDifficulty, newState) {
				upNext.Insert(newState)
			}
		}
	}

	lastCoords := coordinates{len(grid[0]) - 1, len(grid) - 1}
	return coordsDifficulty[serializeCoords(lastCoords)]
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
