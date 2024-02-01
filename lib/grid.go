package lib

import (
	"fmt"
	"strings"
)

func StringToGrid(input string) (grid [][]byte) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	return
}

func GridToString(grid [][]byte) string {
	var lines []string
	for _, line := range grid {
		lines = append(lines, string(line))
	}

	return strings.Join(lines, "\n")
}

func PointInGrid(x, y int, grid [][]byte) bool {
	return IndexInSlice(y, grid) && IndexInSlice(x, grid[y])
}

func PrintGrid(grid [][]byte) {
	for y := range grid {
		fmt.Println(string(grid[y]))
	}
}

// Flips the x and y axis of a grid
func FlipGrid(grid [][]byte) [][]byte {
	cols := len(grid[0])
	rows := len(grid)

	newGrid := CreateGrid(cols, rows, 0)

	// Assign values from the original grid to the new grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newGrid[j][i] = grid[i][j]
		}
	}

	return newGrid
}

func CreateGrid(rows, cols int, defaultValue byte) [][]byte {
	grid := make([][]byte, rows)
	for i := range grid {
		grid[i] = make([]byte, cols)
		for j := range grid[i] {
			grid[i][j] = defaultValue
		}
	}

	return grid
}

func ReverseGrid(grid [][]byte) [][]byte {
	for i := range grid {
		grid[i] = ReverseSlice(grid[i])
	}

	return grid
}

func GridAreEqual(grid1, grid2 [][]byte) bool {
	if len(grid1) != len(grid2) {
		return false
	}

	for i := range grid1 {
		if !EqualSlices(grid1[i], grid2[i]) {
			return false
		}
	}

	return true
}
