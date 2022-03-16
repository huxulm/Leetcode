package maximumproductsubarray

func maxProduct(nums []int) (ans int) {
	n := len(nums)

	dp, ans := [2]int{nums[0], nums[0]}, nums[0]

	for i := 2; i <= n; i++ {
		dp[0], dp[1] = min(nums[i-1], min(nums[i-1]*dp[0], nums[i-1]*dp[1])), max(nums[i-1], max(dp[1]*nums[i-1], dp[0]*nums[i-1]))
		ans = max(ans, dp[1])
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
