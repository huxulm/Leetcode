package maximumsubarray

func maxSubArray(nums []int) int {
	n := len(nums)

	// dp[i] 以nums[i-1]结尾的最大连续和
	// 1<=i<=n
	// 结果: max(dp[1]..dp[n])
	dp := make([]int, n+1)

	maxSum := -1 << 31

	for i := 1; i < n+1; i++ {
		if i == 1 {
			dp[i] = nums[i-1]
			maxSum = max(maxSum, dp[i])
			continue
		}
		dp[i] = max(dp[i-1]+nums[i-1], nums[i-1])
		// 更新答案
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
