package kdiffpairsinanarray

import "sort"

func findPairs(nums []int, k int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := sort.SearchInts(nums[i+1:], k+nums[i])
		if j != n-i-1 && nums[i+j+1] == k+nums[i] {
			ans++
		}
	}
	return
}
