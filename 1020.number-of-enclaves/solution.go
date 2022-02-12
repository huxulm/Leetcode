package numberofenclaves

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numEnclaves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)

	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(i, j int)

	// 深度优先
	dfs = func(i, j int) {
		// 超过边界或被标记过
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || vis[i][j] || grid[i][j] != 1 {
			return
		}
		vis[i][j] = true
		for _, dir := range dirs {
			dfs(dir.x+i, dir.y+j)
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans++
			}
		}
	}

	return ans
}
