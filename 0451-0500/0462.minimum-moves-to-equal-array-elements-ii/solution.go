package minimummovestoequalarrayelementsii

import "sort"

func minMoves2(nums []int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	for _, x := range nums {
		ans += abs(x - nums[n/2])
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
