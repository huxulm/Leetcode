package triangle

func minimumTotal(triangle [][]int) (ans int) {
	n := len(triangle)
	dp := make([][]int, n)
	copy(dp, triangle)

	// 自顶向下计算到达每一层，每个位置的最小路径
	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}
		for j := 0; j < len(dp[i]); j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(dp[i])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}

	ans = dp[n-1][0]
	for i := range dp[n-1] {
		ans = min(ans, dp[n-1][i])
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
