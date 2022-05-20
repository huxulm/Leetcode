package findtheduplicatenumber

import "sort"

// 二分
// 时间 O(NlogN)
// 空间 O(1)
func findDuplicate(nums []int) (ans int) {
	N := len(nums)

	l, r := 1, N-1
	for l < r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < N; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return
}

func findDuplicate1(nums []int) (ans int) {
	N := len(nums)
	// f(x) 先 false 后 true
	return sort.Search(N-1, func(mid int) bool {
		cnt := 0
		for i := range nums {
			if nums[i] <= mid {
				cnt++
			}
			if cnt > mid {
				break
			}
		}
		return cnt > mid
	})
}
