package partitiontokequalsumsubsets

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if n < k || sum%k != 0 {
		return false
	}

	target := sum / k

	var subsets = make([]int, k)

	// 倒排
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == n {
			return true
		}
		for i := range subsets {
			if subsets[i] == target || subsets[i]+nums[idx] > target || (i > 0 && subsets[i] == subsets[i-1]) {
				continue
			}
			subsets[i] += nums[idx]
			if dfs(idx + 1) {
				return true
			}
			subsets[i] -= nums[idx]
		}
		return false
	}
	return dfs(0)
}
