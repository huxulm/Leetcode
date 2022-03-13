package allpathsfromsourcetotarget

func allPathsSourceTarget(g [][]int) (ans [][]int) {
	n := len(g)
	track := []int{0}
	var dfs func(v int)
	dfs = func(v int) {
		if v == n-1 {
			ans = append(ans, append([]int(nil), track...))
			return
		}
		for _, fn := range g[v] {
			track = append(track, fn)
			dfs(fn)
			track = track[:len(track)-1]
		}
	}
	dfs(0)
	return
}
