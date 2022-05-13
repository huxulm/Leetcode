package checkifthereisavalidparenthesesstringpath

func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, (m+n+1)/2)
		}
	}
	var dfs func(i, j, c int) bool
	dfs = func(i, j, c int) bool {
		if c > m-i+n-j-1 {
			return false
		}
		if i == m-1 && j == n-1 { // 终点
			return c == 1 // 终点一定是 ')'
		}
		if vis[i][j][c] {
			return false
		}
		vis[i][j][c] = true
		if grid[i][j] == '(' {
			c++
		} else if c--; c < 0 {
			return false
		}
		return i < m-1 && dfs(i+1, j, c) || j < n-1 && dfs(i, j+1, c)
	}

	if grid[0][0] != '(' || (m+n-1)&1 == 1 || grid[m-1][n-1] != ')' {
		return false
	}
	return dfs(0, 0, 0)
}
