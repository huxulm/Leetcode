package jumpgameii

func jump(nums []int) int {
	var (
		ans    int
		end    int
		maxPos int
	)

	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(nums[i]+i, maxPos)
		if i == end {
			end = maxPos
			ans++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0
	maxPos := 0

	for i := 1; i < n; i++ {
		// 设定总是可以到达最后
		// if i > maxPos {
		// 	break
		// }
		if i+nums[i] <= maxPos {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + 1
			maxPos = i + nums[i]
		}
	}

	return dp[n-1]
}
