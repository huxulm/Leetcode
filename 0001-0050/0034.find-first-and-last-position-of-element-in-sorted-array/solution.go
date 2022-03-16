package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func leftBound(nums []int, target int) (ans int) {
	n := len(nums)
	l, r := 0, n-1
	ans = -1
	for l <= r {
		mid := l + (r-l)>>2
		if nums[mid] == target {
			r = mid - 1
			ans = mid
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return
}

func rightBound(nums []int, target int) (ans int) {
	n := len(nums)
	l, r := 0, n-1
	ans = -1
	for l <= r {
		mid := l + (r-l)>>2
		if nums[mid] == target {
			l = mid
			ans = mid
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return
}

// 宫水三叶 https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/solution/sha-sha-gao-bu-qing-ru-he-ding-yi-er-fen-rrj1/
func searchRange1(nums []int, target int) []int {
	ans := []int{-1, -1}
	n := len(nums)
	if n == 0 {
		return ans
	}
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)>>1 // 注意
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] != target {
		return ans
	} else {
		ans[0] = l

		l, r = 0, n-1
		for l < r {
			mid := l + (r-l+1)>>1 // 注意
			if nums[mid] <= target {
				l = mid
			} else {
				r = mid - 1
			}
		}
		ans[1] = l
		return ans
	}
}
