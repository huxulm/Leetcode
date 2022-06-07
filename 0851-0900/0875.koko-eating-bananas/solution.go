package kokoeatingbananas

import "sort"

func minEatingSpeed(piles []int, h int) int {
	return sort.Search(1e9, func(k int) bool {
		// k++
		tot := 0
		for i := range piles {
			// 向上取整 (a+b-1)/b
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

// 2022/06/07
// 二分
func minEatingSpeed2(piles []int, h int) int {
	return sort.Search(1e9, func(s int) bool {
		s++
		tot := 0
		for _, v := range piles {
			tot += (v + s - 1) / s
		}
		return tot <= h
	})
}

func minEatingSpeed3(piles []int, h int) int {
	var check = func(s int) bool {
		tot := 0
		for i := range piles {
			tot += (piles[i] + s - 1) / s
		}
		return tot <= h
	}
	l, r := 1, -1

	for i := range piles {
		if piles[i] > r {
			r = piles[i]
		}
	}

	for l < r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
