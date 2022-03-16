package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) (res [][]int) {

	sort.Ints(candidates)

	track := []int{}
	var backtrack func(target, idx int)

	backtrack = func(target, idx int) {
		n := len(candidates)
		if target < 0 {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), track...))
			return
		}
		for i := idx; i < n; i++ {
			// 要对同一树层使用过的元素进行跳过
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			if target >= candidates[i] {
				track = append(track, candidates[i])
				backtrack(target-candidates[i], i+1)
				track = (track)[:len(track)-1]
			}
		}
	}

	backtrack(target, 0)

	return
}

func combinationSum2_1(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	n := len(candidates)
	track := []int{}
	var backtrack func(target, sum, idx int)
	backtrack = func(target, sum, idx int) {
		if sum == target {
			ans = append(ans, append([]int(nil), track...))
			return
		}

		if idx == n {
			return
		}

		for i := idx; i < n; i++ {
			// 要对同一树层使用过的元素进行跳过
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			if target >= candidates[i]+sum {
				track = append(track, candidates[i])
				sum += candidates[i]
				backtrack(target, sum, i+1)
				sum -= candidates[i]
				track = track[:len(track)-1]
			}
		}
	}
	backtrack(target, 0, 0)
	return
}
