package maximumcandiesallocatedtokchildren

import "sort"

func maximumCandies(candies []int, k int64) int {
	var check = func(m int) bool {
		var tot int64
		for i := range candies {
			tot += int64(candies[i] / m)
		}
		return tot < k
	}

	l, r := 1, int(1e7)
	for l < r {
		mid := l + (r-l)/2
		if check(mid) { // 不满足要求
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

// https://github.com/EndlessCheng/codeforces-go/blob/dfb9c1e401ac94b77fc8260c10d34d61372b5647/copypasta/sort.go#L92
func maximumCandies1(candies []int, k int64) int {
	return sort.Search(1e7, func(m int) bool {
		var tot int64
		for i := range candies {
			tot += int64(candies[i] / (m + 1))
		}
		return tot < k
	})
}
