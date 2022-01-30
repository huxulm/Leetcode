package minimumpathsum

func minPathSum(grid [][]int) (ans int) {
	ans = 1<<31 - 1
	m, n := len(grid), len(grid[0])
	dfs(&ans, grid, m, n, 0, 0, grid[0][0])
	return
}

// 超时
func dfs(minSum *int, grid [][]int, m, n, i, j, sum int) {
	// 到达终点
	if i == m-1 && j == n-1 {
		*minSum = min(*minSum, sum)
		return
	}
	// 选择向右或向下出发
	if j < n-1 {
		// 选择向右
		sum += grid[i][j+1]
		dfs(minSum, grid, m, n, i, j+1, sum)
		sum -= grid[i][j+1]
	}
	if i < m-1 {
		// 选择向下
		sum += grid[i+1][j]
		dfs(minSum, grid, m, n, i+1, j, sum)
		sum -= grid[i+1][j]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
