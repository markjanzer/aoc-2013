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

func moveAllDirections(position tile, grid [][]byte) (nextPosition []tile) {
	// Up
	if position.row > 0 && grid[position.row-1][position.col] != '#' {
		nextPosition = append(nextPosition, tile{position.row - 1, position.col})
	}
	// Down
	if position.row < len(grid)-1 && grid[position.row+1][position.col] != '#' {
		nextPosition = append(nextPosition, tile{position.row + 1, position.col})
	}
	// Left
	if position.col > 0 && grid[position.row][position.col-1] != '#' {
		nextPosition = append(nextPosition, tile{position.row, position.col - 1})
	}
	// Right
	if position.col < len(grid[0])-1 && grid[position.row][position.col+1] != '#' {
		nextPosition = append(nextPosition, tile{position.row, position.col + 1})
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

func removeDuplicates(queue []tile) []tile {
	cache := map[tile]bool{}
	result := []tile{}
	for _, position := range queue {
		if _, ok := cache[position]; ok {
			continue
		} else {
			cache[position] = true
			result = append(result, position)
		}
	}
	return result
}

func solvePart1(input string, steps int) int {
	grid := lib.StringToGrid(input)
	startingTile := startingTile(grid)

	queue := []tile{startingTile}
	for i := 0; i < steps; i++ {
		nextQueue := []tile{}
		for _, position := range queue {
			nextQueue = append(nextQueue, moveAllDirections(position, grid)...)
		}

		queue = removeDuplicates(nextQueue)
	}

	return len(queue)
}

/*
	Part 2 Notes

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
