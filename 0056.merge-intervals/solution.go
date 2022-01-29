package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		// 按区间左侧值升序排列
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{}
	var left []int
	left = intervals[0]
	for i := 1; i < len(intervals); i++ {
		result, ok := mergeTwo(left, intervals[i])
		if ok {
			left = result[0]
		} else {
			ans = append(ans, left)
			left = intervals[i]
		}
	}
	ans = append(ans, left)
	return ans
}

func mergeTwo(a, b []int) ([][]int, bool) {
	// 需要合并
	if a[1] >= b[0] {
		return [][]int{{a[0], max(a[1], b[1])}}, true
	}
	return nil, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
