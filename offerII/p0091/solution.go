package p0091

func minCost(costs [][]int) int {
	n := len(costs)

	dp := make([][]int, n)
	dp[0] = append(dp[0], costs[0]...)
	for i := 1; i < n; i++ {
		dp[i][0] = costs[i][0] + min(dp[i-1][1], dp[i-1][2])
		dp[i][1] = costs[i][1] + min(dp[i-1][0], dp[i-1][2])
		dp[i][2] = costs[i][2] + min(dp[i-1][0], dp[i-1][1])
	}
	return min(dp[n-1][0], min(dp[n-1][1], dp[n-1][2]))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
