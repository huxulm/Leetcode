package combinations

func combine(n int, k int) [][]int {
	var ans [][]int
	var track = make([]int, k)
	var dfs func(prev, idx int, ans *[][]int)
	dfs = func(prev, idx int, ans *[][]int) {
		if idx == k {
			res := make([]int, k)
			copy(res, track)
			*ans = append(*ans, res)
			return
		}
		// 选择第i个数
		// 选择范围
		for i := prev + 1; i <= n-k+idx; i++ {
			// 选择
			track[idx] = i + 1
			// 选择idx+1个数
			dfs(i, idx+1, ans)
		}
	}
	dfs(-1, 0, &ans)
	return ans
}

// 0 <= idx <= k-1 第idx个数
// prev = idx-1 第idx-1个数
func dfs(n, k, prev, idx int, track []int, ans *[][]int) {
	if idx == k {
		res := make([]int, k)
		copy(res, track)
		*ans = append(*ans, res)
		return
	}
	// 选择第i个数
	// 选择范围
	for i := prev + 1; i <= n-k+idx; i++ {
		// 选择
		track[idx] = i + 1
		// 选择idx+1个数
		dfs(n, k, i, idx+1, track, ans)
	}
}
