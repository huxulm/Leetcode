package numberofprovinces

import "container/list"

func findCircleNum(c [][]int) (ans int) {
	g := make([][]int, len(c))

	for i := range c {
		for j := range c[i] {
			if len(g[i]) == 0 {
				g[i] = []int{}
			}
			if len(g[j]) == 0 {
				g[j] = []int{}
			}
			if c[i][j] == 1 {
				// 有向图
				g[i] = append(g[i], j)
			}
		}
	}

	var vis = make([]bool, len(c))
	var bfs func(v int)
	bfs = func(v int) {
		vis[v] = true
		q := list.New()
		q.PushBack(v)
		for q.Len() > 0 {
			s := q.Remove(q.Front()).(int)
			for _, fs := range g[s] {
				if !vis[fs] {
					vis[fs] = true
					q.PushBack(fs)
				}
			}
		}
	}

	for v := range g {
		if !vis[v] {
			ans++
			bfs(v)
		}
	}

	return
}
