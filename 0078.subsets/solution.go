package subsets

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	// nums 的 0~n个元素的组合
	for i := 0; i <= n; i++ {
		combine(nums, n, i, &ans)
	}
	return
}

func combine(nums []int, n int, k int, ans *[][]int) {
	var track = make([]int, k)
	var dfs func(prev, idx int)
	dfs = func(prev, idx int) {
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
			track[idx] = nums[i]
			// 选择idx+1个数
			dfs(i, idx+1)
		}
	}
	dfs(-1, 0)
	return
}

// LC official
// 迭代
func subsets1(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

// 递归(dfs)
func subsets2(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
