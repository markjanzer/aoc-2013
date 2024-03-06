package main

import (
	"advent-of-code-2023/lib"
	"fmt"
)

const SmallTestString string = ``

const TestString string = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Alright, here's what we do.
	Create a grid with all of the tiles
	Create a function that returns possible moves from a given point.
	Add all of those positions to a queue for the next move, but don't add redundant moves
	Do this 64 times.

*/

type tile struct {
	row, col int
}

func moveAllDirections(position tile, grid [][]byte, oddity string, tileCache *map[tile]string) (nextPosition []tile) {
	// Up
	nextTile := tile{position.row - 1, position.col}
	if nextTile.row >= 0 && grid[nextTile.row][nextTile.col] != '#' {
		if _, ok := (*tileCache)[nextTile]; !ok {
			(*tileCache)[nextTile] = oddity
			nextPosition = append(nextPosition, nextTile)
		}
	}

	// Down
	nextTile = tile{position.row + 1, position.col}
	if nextTile.row < len(grid) && grid[nextTile.row][nextTile.col] != '#' {
		if _, ok := (*tileCache)[nextTile]; !ok {
			(*tileCache)[nextTile] = oddity
			nextPosition = append(nextPosition, nextTile)
		}
	}

	// Left
	nextTile = tile{position.row, position.col - 1}
	if nextTile.col >= 0 && grid[nextTile.row][nextTile.col] != '#' {
		if _, ok := (*tileCache)[nextTile]; !ok {
			(*tileCache)[nextTile] = oddity
			nextPosition = append(nextPosition, nextTile)
		}
	}

	// Right
	nextTile = tile{position.row, position.col + 1}
	if nextTile.col < len(grid[0]) && grid[nextTile.row][nextTile.col] != '#' {
		if _, ok := (*tileCache)[nextTile]; !ok {
			(*tileCache)[nextTile] = oddity
			nextPosition = append(nextPosition, nextTile)
		}
	}

	return
}

func startingTile(grid [][]byte) tile {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				return tile{row, col}
			}
		}
	}
	panic("Starting tile not found")
}

func solvePart1(input string, steps int) int {
	grid := lib.StringToGrid(input)
	startingTile := startingTile(grid)
	tileCache := map[tile]string{}

	queue := []tile{startingTile}
	for i := 0; i < steps; i++ {
		oddity := evenOrOdd(i + 1)
		nextQueue := []tile{}
		for _, position := range queue {
			nextQueue = append(nextQueue, moveAllDirections(position, grid, oddity, &tileCache)...)
		}

		queue = nextQueue
	}

	desiredOddity := evenOrOdd(steps)
	count := 0
	for _, oddity := range tileCache {
		if oddity == desiredOddity {
			count++
		}
	}
	return count
}

func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}

/*
	Part 2 Notes

	Okay, we can start by making part 1 more efficient.
	Instead of creating a new tile for every possible move, we can keep track of whether
	the tile was reached in an even or odd number of steps.
	For now we can have a map of all of all of the tiles and determine if it's even or odd.

*/

type coords struct {
	x, y int
}

type gridCacheData struct {
	coords coords
	step   int
}

// It keeps track of the coordinates of the first tile accessed in a given grid
// One grid to the left is -1, 0 for coords, and then the coords are used very
// much like tile
var newGridCache = make(map[coords]gridCacheData)

func canMove(position tile, grid [][]byte) bool {
	rows := len(grid)
	cols := len(grid[0])

	gridX := position.row / rows
	gridY := position.col / cols

	gridRow := (rows + (position.row % rows)) % rows
	gridCol := (cols + (position.col % cols)) % cols

	if grid[gridRow][gridCol] == '#' {
		return false
	}

	gridCoords := coords{gridX, gridY}
	if _, ok := newGridCache[gridCoords]; !ok {
		newTileCoords := coords{gridRow, gridCol}
		fmt.Println(gridCoords, newTileCoords, step)
		newGridCache[gridCoords] = gridCacheData{newTileCoords, step}
	}

	return true
}

func moveAllDirections2(position tile, grid [][]byte, oddity string, tileCache *map[tile]string) (nextPosition []tile) {
	var nextTile tile

	allDirections := []tile{
		{position.row - 1, position.col},
		{position.row + 1, position.col},
		{position.row, position.col - 1},
		{position.row, position.col + 1},
	}

	for _, nextTile = range allDirections {
		if canMove(nextTile, grid) {
			if _, ok := (*tileCache)[nextTile]; !ok {
				(*tileCache)[nextTile] = oddity
				nextPosition = append(nextPosition, nextTile)
			}
		}
	}

	return
}

var step = 0

/*
Part 2 Notes, discussing the solution

I think I know how to solve part 2, but it looks like it'll take a bit of time,
and I want to work on other things atm. Here's my proposed solution:

Since part 2 can go forever, assume there is not only coordinates for each input grid, but
also a grid of those inputs--if you start at 0,0 in the first grid (I'll call it a square)
and then move to the left then you'll end up at 65,0 of square -1,0

Pretty quickly, we start entering new grids at a consistently increasing rate.
We see a repeating pattern every 131 steps starting around 195
In the 131 steps starting with 195 we enter 8 new squares
In the 131 steps starting with 326 we enter 12 new squares
In the 131 steps starting with 457 we enter 16 new squares
etc.

There also is predictability on what step they are entering and which square they
are starting with.

We can figure how long it takes to fill a square from each starting point.
We can also figure out how many tiles will be counted given an even or odd step.
This gets us very close, because we will know how many even + odd squares will be full by
26501365 steps.
The last part is to figure out how far into each of the currently expanding squares
this number is, and then multiply that value for the number of squares the pattern will
be filling by that step.
*/

func solvePart2(input string, steps int) int {
	grid := lib.StringToGrid(input)
	startingTile := startingTile(grid)
	tileCache := map[tile]string{}

	queue := []tile{startingTile}
	for i := 0; i < steps; i++ {
		step = i
		oddity := evenOrOdd(i + 1)
		nextQueue := []tile{}
		for _, position := range queue {
			nextQueue = append(nextQueue, moveAllDirections2(position, grid, oddity, &tileCache)...)
		}

		queue = nextQueue
	}

	desiredOddity := evenOrOdd(steps)
	count := 0
	for _, oddity := range tileCache {
		if oddity == desiredOddity {
			count++
		}
	}
	return count
}

func main() {
	lib.AssertEqual(16, solvePart1(TestString, 6))
	lib.AssertEqual(1594, solvePart2(TestString, 50))
	// lib.AssertEqual(6536, solvePart2(TestString, 100))
	// lib.AssertEqual(167004, solvePart2(TestString, 500))
	// lib.AssertEqual(668697, solvePart2(TestString, 1000))
	// lib.AssertEqual(16733044, solvePart2(TestString, 5000))

	// fmt.Println(newGridCache)

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString, 64)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString, 26501365)
	// result2 := solvePart2(dataString, 1000)
	// fmt.Println(result2)
}
