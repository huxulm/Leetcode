package subsetsii

import "sort"

func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	var track = []int{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			ans = append(ans, append([]int(nil), track...))
			return
		}

		repeats := []int{nums[cur]}
		for cur < n-1 && nums[cur] == nums[cur+1] {
			cur++
			repeats = append(repeats, nums[cur])
		}

		// 选repeats
		for i := range repeats {
			track = append(track, repeats[:i+1]...)
			dfs(cur + 1)
			track = track[:len(track)-i-1]
		}

		// 不选
		dfs(cur + 1)
	}
	return
}
