package searchinrotatedsortedarray

func search(nums []int, target int) int {

	n := len(nums)
	lo, hi := 0, n-1

	if n < 1 {
		return -1
	}

	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	for lo <= hi && hi < len(nums) {

		mid := (lo + hi) >> 1

		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
