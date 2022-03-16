package arithmeticslices

// 输入：nums = [1,2,3,4]
// 输出：3
// 解释：nums 中有三个子等差数组：[1, 2, 3]、[2, 3, 4] 和 [1,2,3,4] 自身。
func numberOfArithmeticSlices(nums []int) (ans int) {
	n := len(nums)
	if n == 1 {
		return 0
	}

	d, t := nums[0]-nums[1], 0
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d = nums[i-1] - nums[i]
			t = 0
		}
		ans += t
	}
	return
}
