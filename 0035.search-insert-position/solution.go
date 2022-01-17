package searchinsertposition

func searchInsert(nums []int, target int) int {
	return left_bound(nums, target)
}

func left_bound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			r = mid
		} else if nums[mid] < target {
			l = mid + 1 // 注意
		} else {
			r = mid // 注意
		}
	}
	return l
}
