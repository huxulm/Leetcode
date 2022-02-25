package houserobber

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	// 对于前nums[0...i], 第i个数取不/取能得到的最大值
	dp := make([]int, n+1)
	dp[1] = nums[0]
	dp[2] = max(dp[1], nums[1])
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
