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

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(16, solvePart1(TestString, 6))
	// lib.AssertEqual(1, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	dataString := lib.GetDataString(DataFile)
	result1 := solvePart1(dataString, 64)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
