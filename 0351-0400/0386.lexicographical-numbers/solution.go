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

func lexicalOrder1(n int) (ans []int) {
	ans = make([]int, n)

	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return
}
