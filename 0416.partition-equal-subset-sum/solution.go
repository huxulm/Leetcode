package partitionequalsubsetsum

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	sum /= 2

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] >= 0 {
				// 不选或选nums[i]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][sum]
}
