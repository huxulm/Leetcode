package heightchecker

import "sort"

func heightChecker(heights []int) (ans int) {
	n := len(heights)
	expected := make([]int, n)
	copy(expected, heights)
	sort.Ints(expected)
	for i := range expected {
		if expected[i] != heights[i] {
			ans++
		}
	}
	return
}
