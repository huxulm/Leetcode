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

func permuteUnique1(nums []int) (ans [][]int) {
	sort.Ints(nums)

	n := len(nums)
	track := make([]int, n)
	used := make([]bool, n)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), track...))
			return
		}

		for i, x := range nums {
			if used[i] || (i > 0 && x == nums[i-1] && !used[i-1]) {
				continue
			}
			// 选择 idx
			track[idx] = x
			// 标记 idx
			used[i] = true
			// 选择 idx+1
			backtrack(idx + 1)
			// 撤销 idx 标记
			used[i] = false
		}
	}

	backtrack(0)
	return
}
