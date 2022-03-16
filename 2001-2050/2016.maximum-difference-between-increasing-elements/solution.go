package maximumdifferencebetweenincreasingelements

func maximumDifference(nums []int) (ans int) {
	n := len(nums)
	pre_min := nums[0]
	ans = -1
	for i := 1; i < n; i++ {
		if nums[i] > pre_min {
			ans = max(ans, nums[i]-pre_min)
		} else {
			pre_min = nums[i]
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
