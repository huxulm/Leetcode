package minimumtimetocompletetrips

import "sort"

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
