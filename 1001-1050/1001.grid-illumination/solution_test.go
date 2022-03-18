package gridillumination

import (
	"fmt"
	"testing"
)

func TestOpenLamp(t *testing.T) {
	// grid := genGrid(6)
	lamps := [][]int{{2, 5}, {4, 2}, {0, 3}, {0, 5}, {1, 4}, {4, 2}, {3, 3}, {1, 0}}
	querys := [][]int{{4, 3}, {3, 1}, {5, 3}, {0, 5}, {4, 4}, {3, 3}}
	fmt.Println(gridIllumination(6, lamps, querys))
}

func genGrid(n int) [][]int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	return grid
}

func PrintGrid(grid [][]int) {
	for i := range grid {
		fmt.Printf("%v\n", grid[i])
	}
}
