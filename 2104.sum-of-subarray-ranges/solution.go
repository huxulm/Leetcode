package sumofsubarrayranges

// 暴力 O(n^3)
func subArrayRanges(nums []int) (ans int64) {
	n := len(nums)
	// 子数组长度
	for i := 1; i <= n; i++ {
		// 子数组范围
		for j := 0; j <= n-i; j++ {
			v1, v2 := MinMax(nums[j : j+i])
			ans += int64(v2 - v1)
		}
	}
	return
}

func MinMax(nums []int) (int, int) {
	min, max := 1<<31-1, -1<<31
	for _, x := range nums {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	return min, max
}

// 方法二: 先枚举左边界，再枚举右边界
func subArrayRanges1(nums []int) (ans int64) {
	n := len(nums)
	for i := 0; i < n; i++ {
		min, max := nums[i], nums[i]
		for j := i + 1; j < n; j++ {
			if min > nums[j] {
				min = nums[j]
			}
			if max < nums[j] {
				max = nums[j]
			}
			ans += int64(max - min)
		}
	}
	return
}
