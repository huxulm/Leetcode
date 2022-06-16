package findkthsmallestpairdistance

import "sort"

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	// 答案区间 [nums[0], nums[n-1]]
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt := 0
		for j, num := range nums {
			// 查找大于等于 nums[j]−mid 的最小值的下标 i
			i := sort.SearchInts(nums[:j], num-mid)
			// 右端点为 j 且距离小于等于 mid 的数对数目为 j - i
			cnt += j - i
		}
		return cnt >= k
	})
}

func smallestDistancePair1(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt, i := 0, 0
		for j, num := range nums {
			for num-nums[i] > mid {
				i++
			}
			cnt += j - i
		}
		return cnt >= k
	})
}
