package uniquepaths

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// i = 0
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// j = 0
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths1(m, n int) int {
	return dp(m-1, n-1)
}

// 定义：从 (0, 0) 到 (x, y) 有 dp(x, y) 条路径
func dp(x, y int) int {
	if x == 0 && y == 0 {
		return 1
	}
	if x < 0 || y < 0 {
		return 0
	}
	// 状态转移方程：
	// 到达 (x, y) 的路径数等于到达 (x - 1, y) 和 (x, y - 1) 路径数之和
	return dp(x-1, y) + dp(x, y-1)
}

// 动态规划
func uniquePaths2(m int, n int) int {
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[m-1]
}
