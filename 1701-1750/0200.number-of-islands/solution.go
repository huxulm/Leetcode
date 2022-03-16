package numberofislands

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return false
		}
		grid[i][j] = 0
		for _, d := range dirs {
			if x, y := i+d[0], j+d[1]; x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
				dfs(x, y)
			}
		}
		return true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j) {
				ans++
			}
		}
	}
	return
}
