package sum

import "sort"

func twoSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	lo, hi := 0, len(nums)-1
	ans := make([][]int, 0)
	for lo < hi {
		left, right := nums[lo], nums[hi]
		sum := left + right
		if sum == target {
			ans = append(ans, []int{left, right})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	ans := make([][]int, 0)

	return ans
}
