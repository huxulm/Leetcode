package minimumheighttrees

// 方法一: BFS
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	parents := make([]int, n)
	bfs := func(start int) (x int) {
		vis := make([]bool, n)
		vis[start] = true
		q := []int{start}
		for len(q) > 0 {
			x, q = q[0], q[1:]
			for _, y := range g[x] {
				if !vis[y] {
					vis[y] = true
					parents[y] = x
					q = append(q, y)
				}
			}
		}
		return
	}
	x := bfs(0) // 找到与节点 0 最远的节点 x
	y := bfs(x) // 找到与节点 x 最远的节点 y

	path := []int{}
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}
}

func findMinHeightTrees1(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	graph := make([][]int, n)
	degrees := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
		degrees[u]++
		degrees[v]++
	}
	var q []int
	for i := 0; i < n; i++ {
		if degrees[i] == 1 {
			q = append(q, i)
		}
	}
	var ans []int
	for len(q) > 0 {
		ans = q
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[i]
			for _, neighbor := range graph[node] {
				degrees[neighbor]--
				if degrees[neighbor] == 1 {
					q = append(q, neighbor)
				}
			}
		}
		q = q[size:]
	}
	return ans
}
