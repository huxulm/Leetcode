package pacificatlanticwaterflow

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {

	m, n := len(heights), len(heights[0])

	pacific, atlantic := make([][]bool, m), make([][]bool, m)
	for i := range pacific {
		pacific[i], atlantic[i] = make([]bool, n), make([]bool, n)
	}

	var dfs func(r, c int, vis [][]bool)
	dfs = func(r, c int, vis [][]bool) {
		if vis[r][c] {
			return
		}

		vis[r][c] = true

		for _, dir := range dirs {
			x, y := r+dir[0], c+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && heights[x][y] >= heights[r][c] {
				dfs(x, y, vis)
			}
		}
	}

	for i := 0; i < n; i++ {
		dfs(0, i, pacific)
	}
	for j := 0; j < m; j++ {
		dfs(j, 0, pacific)
	}

	for i := 0; i < n; i++ {
		dfs(m-1, i, atlantic)
	}
	for j := 0; j < m; j++ {
		dfs(j, n-1, atlantic)
	}

	for i := range pacific {
		for j := range pacific[i] {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
