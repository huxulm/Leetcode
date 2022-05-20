package findrightinterval

import "sort"

func findRightInterval(intervals [][]int) (ans []int) {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}

	// O(nlog(n))
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	n := len(intervals)
	ans = make([]int, n)
	// O(nlog(n))
	for i := range intervals {
		pos := sort.Search(n, func(x int) bool { return intervals[x][0] >= intervals[i][1] })
		if pos == n {
			ans[intervals[i][2]] = -1
		} else {
			ans[intervals[i][2]] = intervals[pos][2]
		}
	}
	return
}
