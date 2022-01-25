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
