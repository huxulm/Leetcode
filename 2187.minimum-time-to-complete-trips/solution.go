package minimumtimetocompletetrips

import (
	"math"
	"sort"
)

// 方法一: 库函数二分判定
func minimumTime(time []int, totalTrips int) int64 {
	return int64(sort.Search(totalTrips*1e7, func(x int) bool {
		tot := 0
		for i := range time {
			tot += x / time[i]
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}

// 方法二: 手写二分判定
func minimumTime1(time []int, totalTrips int) int64 {
	l, r := 1, totalTrips*1e7
	var check = func(t int) bool {
		tot := 0
		for i := range time {
			tot += t / time[i]
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}
	// 二分查找寻找满足要求的最小的 t
	for l < r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}

func minimumTime2(time []int, totalTrips int) int64 {
	return lowerSearch(1, math.MaxInt64, func(i int64) bool {
		r := int64(0)
		for j := range time {
			if i >= int64(time[j]) {
				r += i / int64(time[j])
			}
			if r >= int64(totalTrips) {
				return true
			}
		}
		return r >= int64(totalTrips)
	})
}

func lowerSearch(l, r int64, f func(int64) bool) int64 {
	for l < r {
		h := int64(uint(r+l) >> 1)
		if !f(h) {
			l = h + 1 // 保留 f(l-1) = false
		} else {
			r = h // 保留 f(r) = true
		}
	}
	return l
}
