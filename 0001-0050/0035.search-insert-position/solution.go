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

//「在一个有序数组中找第一个大于等于 \textit{target}target 的下标」。
func searchInsert1(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (r + l) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
