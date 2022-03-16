package minimumpathsum

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	return dp(grid, m-1, n-1, make(map[string]int))
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

// 自顶向下
func dp(grid [][]int, x, y int, memo map[string]int) (ans int) {
	key := fmt.Sprintf("%d,%d", x, y)

	if v, ok := memo[key]; ok {
		return v
	}

	// base case
	if x == 0 && y == 0 {
		ans = grid[x][y]
	} else if x == 0 && y > 0 {
		ans = grid[x][y] + dp(grid, x, y-1, memo)
	} else if y == 0 && x > 0 {
		ans = grid[x][y] + dp(grid, x-1, y, memo)
	} else {
		ans = grid[x][y] + min(dp(grid, x, y-1, memo), dp(grid, x-1, y, memo))
	}
	memo[key] = ans
	return ans
}

func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for j := 0; j < n; j++ {
		if j == 0 {
			dp[0][j] = grid[0][j]
		} else {
			dp[0][j] = grid[0][j] + dp[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[n-1][m-1]
}
