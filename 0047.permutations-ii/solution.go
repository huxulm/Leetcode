package permutationsii

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	var backtrack func(nums []int, marked []bool, track []int, intdex int)

	backtrack = func(nums []int, marked []bool, track []int, index int) {
		if len(nums) == index {
			dst := make([]int, len(nums))
			copy(dst, track)
			res = append(res, dst)
			return
		}
		for i := 0; i < len(nums); i++ {
			if marked[i] || (i > 0 && nums[i] == nums[i-1] && !marked[i-1]) {
				continue
			}
			// 选择
			track[index] = nums[i]
			marked[i] = true
			backtrack(nums, marked, track, index+1)
			// 取消选择
			marked[i] = false
		}
	}

	n := len(nums)
	backtrack(nums, make([]bool, n), make([]int, n), 0)

	return res
}
