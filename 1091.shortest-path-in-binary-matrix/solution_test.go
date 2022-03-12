package shortestpathinbinarymatrix

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}}
	shortestPathBinaryMatrix(grid)
}
