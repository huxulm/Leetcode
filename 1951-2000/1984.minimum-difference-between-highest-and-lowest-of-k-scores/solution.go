package minimumdifferencebetweenhighestandlowestofkscores

import (
	"sort"
)

// sliding window
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)

	if k > n {
		return 0
	}

	minDiff := 1<<31 - 1

	for right := k - 1; right < n; right++ {
		if diff := nums[right] - nums[right+1-k]; diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
