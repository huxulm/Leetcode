package coinchange

func coinChange(coins []int, amount int) int {
	// F(n) = 1+min{F(n-coins[0], ... , F(n-coins[i]))}  (F(n) != -1)
	dp := make([]int, amount+1)
	// dp[0] = 0
	for i := 1; i <= amount; i++ {
		ans := 1<<31 - 1
		for j := range coins {
			if i-coins[j] >= 0 && dp[i-coins[j]] > 0 {
				ans = min(ans, dp[i-coins[j]])
			}
		}
		if ans == (1<<31 - 1) {
			dp[i] = -1
		} else {
			dp[i] = ans + 1
		}

	}

	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
