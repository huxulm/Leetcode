package combinationsum

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// 升序
	sort.Ints(candidates)

	res := make([][]int, 0)
	track := []int{}

	backtrack1(&res, candidates, target, &track, 0)

	return res
}

func backtrack(res *[][]int, canditates []int, target int, track *[]int, idx int) {
	if idx == len(canditates) {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int(nil), *track...))
		return
	}

	// 选择当前
	if target >= canditates[idx] {
		// 选择
		*track = append(*track, canditates[idx])
		backtrack(res, canditates, target-canditates[idx], track, idx)
		// 撤销选择
		*track = (*track)[:len(*track)-1]
	}

	// 直接跳过当前
	backtrack(res, canditates, target, track, idx+1)
}

func backtrack1(res *[][]int, canditates []int, target int, track *[]int, idx int) {
	if idx == len(canditates) {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int(nil), *track...))
		return
	}

	// 搜索起点, 已经考虑过的数不再考虑
	for i := idx; i < len(canditates); i++ {
		if target >= canditates[i] {
			*track = append(*track, canditates[i])
			backtrack1(res, canditates, target-canditates[i], track, i)
			*track = (*track)[:len(*track)-1]
		}
	}
}

func combinationSum1(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	n := len(candidates)

	if candidates[0] > target {
		return
	}

	track := []int(nil)
	var backtrack func(t, idx int)
	backtrack = func(t, idx int) {

		if t <= 0 || idx == n {
			if t == 0 {
				ans = append(ans, append([]int(nil), track...))
			}
			return
		}

		// 选择当前
		if t >= candidates[idx] {
			track = append(track, candidates[idx])
			backtrack(t-candidates[idx], idx)
			track = track[:len(track)-1]
		}

		// 不选当前
		backtrack(t, idx+1)
	}

	backtrack(target, 0)
	return
}
