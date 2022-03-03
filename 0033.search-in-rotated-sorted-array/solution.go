package searchinrotatedsortedarray

// 线性查找O(n)
func search0(nums []int, target int) int {
	for i, x := range nums {
		if x == target {
			return i
		}
	}
	return -1
}

func search(nums []int, target int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, n-1

	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[0] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[0] {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

func search1(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			// 有序部分下界
			if target >= nums[0] && target < nums[mid] { // 正常半区
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[0] {
			// 有序部分上界
			if target > nums[mid] && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// 方法三: 先找分割点 再去指定区域查找
func search2(nums []int, target int) (ans int) {
	n := len(nums)
	idx := 0
	for i := range nums[:n-1] {
		if nums[i] > nums[i+1] {
			idx = i
			break
		}
	}

	ans = find(nums, 0, idx, target)
	if ans != -1 {
		return
	}
	ans = find(nums, idx+1, n-1, target)
	return
}

func find(nums []int, lo, hi, target int) int {
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		}
	}
	return -1
}
