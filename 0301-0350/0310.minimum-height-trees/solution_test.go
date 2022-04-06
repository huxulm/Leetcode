package minimumheighttrees

import "testing"

func TestFindMinHeightTrees(t *testing.T) {
	findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
}
