package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	if obstacleGrid[0][0] == 1 {
		dp[1][1] = 0
	} else {
		dp[0][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// base case
			if i == 1 && j == 1 {
				continue
			}
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
