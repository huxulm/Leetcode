package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	n := len(prices)
	// base case: dp[-1][0] = 0, dp[-1][1] = -infinity
	dp_i_0 := 0
	dp_i_1 := -1 << 31
	for i := 0; i < n; i++ {
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		// dp[i][1] = max(dp[i-1][1], -prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
