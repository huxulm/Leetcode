package lexicographicalnumbers

func lexicalOrder(n int) (ans []int) {
	ans = make([]int, n)
	idx := 0

	var dfs func(cur int)
	dfs = func(cur int) {
		if cur > n {
			return
		}
		ans[idx] = cur
		idx++
		for i := 0; i <= 9; i++ {
			dfs(cur*10 + i)
		}
	}

	for i := 1; i <= 9; i++ {
		dfs(i)
	}
	return
}
