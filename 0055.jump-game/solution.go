package jumpgame

func canJump(nums []int) bool {
	maxPos := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		// 无法到达i+1
		if nums[i] == 0 && maxPos <= i {
			break
		}
		maxPos = max(maxPos, i+nums[i])
	}
	return maxPos >= n-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
