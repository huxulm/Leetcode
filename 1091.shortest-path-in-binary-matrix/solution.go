package shortestpathinbinarymatrix

var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func shortestPathBinaryMatrix(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	if grid[0][0] != 0 {
		return -1
	}

	q := [][]int{{0, 0}}
	grid[0][0] = 1
	for len(q) > 0 {
		ans++
		tmp := q
		q = nil
		for _, p := range tmp {
			if p[0] == m-1 && p[1] == n-1 {
				return
			}
			for i := range dirs {
				if x, y := p[0]+dirs[i][0], p[1]+dirs[i][1]; x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0 {
					grid[x][y] = 1
					q = append(q, []int{x, y})
				}
			}
		}
	}

	ans = -1
	return
}
