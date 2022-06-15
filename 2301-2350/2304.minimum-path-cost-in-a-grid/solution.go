package minimumpathcostinagrid

func minPathCost(g [][]int, co [][]int) (ans int) {
	m, n := len(g), len(g[0])

	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}

	// 第 i=0 行 花费
	copy(f[0], g[0])

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			f[i][j] = f[i-1][0] + co[g[i-1][0]][j] + g[i][j]
			for k := 0; k < n; k++ {
				f[i][j] = min(f[i][j], f[i-1][k]+co[g[i-1][k]][j]+g[i][j])
			}
		}
	}

	ans = 1<<31 - 1
	for i := range f[m-1] {
		ans = min(f[m-1][i], ans)
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
