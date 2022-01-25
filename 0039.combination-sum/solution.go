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
