package thetimewhenthenetworkbecomesidle

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
	n := len(patience)
	g := make([][]int, n)

	for i := range edges {
		x, y := edges[i][0], edges[i][1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	q := []int{0}
	vis := make([]bool, n)
	vis[0] = true

	for dist := 1; q != nil; dist++ {
		tmp := q
		q = nil

		for _, x := range tmp {
			for _, v := range g[x] {
				if vis[v] {
					continue
				}
				vis[v] = true
				q = append(q, v)
				ans = max(ans, (dist*2-1)/patience[v]*patience[v]+dist*2+1)
			}
		}
	}

	return
}
