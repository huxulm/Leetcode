package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)

	left, right := newInterval[0], newInterval[1]

	merged := false

	ans := [][]int{}

	for i := 0; i < n; i++ {

		if intervals[i][0] > right {
			// 新区间在左侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, intervals[i])
		} else if intervals[i][1] < left {
			// 新区间在右侧且无交集
			ans = append(ans, intervals[i])
		} else {
			left, right = min(left, intervals[i][0]), max(right, intervals[i][1])
		}

	}

	if !merged {
		ans = append(ans, []int{left, right})
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
