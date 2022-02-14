package containsduplicate

import "sort"

func containsDuplicate(nums []int) bool {
	rec := map[int]bool{}

	for i := range nums {
		if !rec[nums[i]] {
			rec[nums[i]] = true
		} else {
			return true
		}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
