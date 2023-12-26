package lib

import (
	"strings"
)

func StringToGrid(input string) (grid [][]byte) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	return
}

func PointInGrid(x, y int, grid [][]byte) bool {
	return IndexInSlice(y, grid) && IndexInSlice(x, grid[y])
}
