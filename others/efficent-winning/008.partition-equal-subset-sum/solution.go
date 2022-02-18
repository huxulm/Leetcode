package partitionequalsubsetsum

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	N := len(nums)
	M := sum / 2

	// 背包问题，选前[i]个数能否凑出和为M
	dp := make([][]bool, N+1)

	for i := range dp {
		dp[i] = make([]bool, M+1)
		dp[i][0] = true
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {

			if j-nums[i-1] >= 0 {
				// 不选或选nums[i]凑成和为j
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else { // 背包容量不足
				dp[i][j] = false
			}

		}
	}

	return dp[N][M]
}
