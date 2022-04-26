package kokoeatingbananas

import "sort"

func minEatingSpeed(piles []int, h int) int {

	// f(x) 先 true 后 false，返回最小的
	// 使得 f(x) 为 true 的答案
	return sort.Search(1e9, func(k int) bool {
		// k++
		tot := 0
		for i := range piles {
			// 向上取整 a+b-1/b
			tot += (piles[i] + k) / (k + 1)
		}
		return tot <= h
	}) + 1
}

func minEatingSpeed1(piles []int, h int) int {
	var check = func(k int) bool {
		tot := 0
		for i := range piles {
			tot += (piles[i] + k - 1) / k
		}
		return tot > h
	}

	l, r := 1, -1

	for i := range piles {
		if piles[i] > r {
			r = piles[i]
		}
	}

	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
