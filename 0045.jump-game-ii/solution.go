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
