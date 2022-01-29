package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		// 按区间左侧值升序排列
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{}
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		if last[1] >= intervals[i][0] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
