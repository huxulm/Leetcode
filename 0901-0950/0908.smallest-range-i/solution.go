package smallestrangei

func smallestRangeI(nums []int, k int) int {
	a, b := nums[0], nums[0]
	for i := range nums {
		if nums[i] < a {
			a = nums[i]
		} else if nums[i] > b {
			b = nums[i]
		}
	}
	// a <= s/n <= b
	if b-a < 2*k {
		return 0
	} else {
		return b - a - 2*k
	}
}
