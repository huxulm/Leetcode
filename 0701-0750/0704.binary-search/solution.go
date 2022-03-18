package binarysearch

func search(nums []int, target int) int {
	return binarysearch_iterative(nums, target)
}

func binarysearch_iterative(nums []int, target int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	l, r := 0, n-1

	for l <= r {
		mid := (l + r) >> 1

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
