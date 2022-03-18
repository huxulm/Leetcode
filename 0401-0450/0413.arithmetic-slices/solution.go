package arithmeticslices

func numberOfArithmeticSlices(nums []int) (ans int) {
	n := len(nums)
	if n < 3 {
		return 0
	}

	d, cnt := nums[1]-nums[0], 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == d {
			cnt++
		} else {
			d = nums[i] - nums[i-1]
			cnt = 0
		}
		ans += cnt
	}

	return
}
