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

func combine1(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}
