package numberofwaystoreconstructatree

import (
	"math"
	"sort"
)

// LC official
func checkWays(pairs [][]int) int {
	adj := map[int]map[int]bool{}
	for _, p := range pairs {
		x, y := p[0], p[1]
		if adj[x] == nil {
			adj[x] = map[int]bool{}
		}
		adj[x][y] = true
		if adj[y] == nil {
			adj[y] = map[int]bool{}
		}
		adj[y][x] = true
	}

	// 检测是否存在根节点
	root := -1
	for node, neighbours := range adj {
		if len(neighbours) == len(adj)-1 {
			root = node
			break
		}
	}
	if root == -1 {
		return 0
	}

	ans := 1
	for node, neighbours := range adj {
		if node == root {
			continue
		}

		currDegree := len(neighbours)
		parent := -1
		parentDegree := math.MaxInt32
		// 根据 degree 的大小找到 node 的父节点 parent
		for neighbour := range neighbours {
			if len(adj[neighbour]) < parentDegree && len(adj[neighbour]) >= currDegree {
				parent = neighbour
				parentDegree = len(adj[neighbour])
			}
		}
		if parent == -1 {
			return 0
		}
		// 检测 neighbours 是否为 adj[parent] 的子集
		for neighbour := range neighbours {
			if neighbour != parent && !adj[parent][neighbour] {
				return 0
			}
		}

		if parentDegree == currDegree {
			ans = 2
		}
	}
	return ans
}

type unordered_set struct {
	m   map[int]struct{}
	set []int
}

func (s unordered_set) Insert(val int) {
	if _, ok := s.m[val]; !ok {
		s.m[val] = struct{}{}
		s.set = append(s.set, val)
	}
}

func (s unordered_set) First() int {
	if len(s.set) == 0 {
		return -1
	}
	return s.set[0]
}

func (s unordered_set) End() int {
	if len(s.set) == 0 {
		return -1
	}
	return s.set[len(s.set)-1]
}

func (s unordered_set) Size() int {
	return len(s.set)
}

func (s unordered_set) Find(v int) int {
	for i := range s.set {
		if s.set[i] == v {
			return i
		}
	}
	return len(s.set) - 1
}

func checkWays1(pairs [][]int) int {
	ways := 1
	g := map[int]unordered_set{}
	t := map[int][]int{}
	vis := map[int]bool{}

	var checkTree func(root, depth int) int
	checkTree = func(root, depth int) int {
		if vis[root] {
			ways = 0
			return 0
		}
		vis[root] = true
		cnt := 0
		for _, c := range t[root] {
			if ways > 0 {
				cnt += checkTree(c, depth+1)
			}
		}

		// 深度（根节点到子节点的路径）+子节点数量 = 祖先关系数量
		if cnt+depth != g[root].Size() {
			ways = 0
			return 0
		}

		return cnt + 1
	}

	if len(pairs) == 1 {
		return 2
	}

	for _, p := range pairs {
		if _, ok := g[p[0]]; !ok {
			g[p[0]] = unordered_set{
				m:   make(map[int]struct{}),
				set: []int{},
			}
		}
		if _, ok := g[p[1]]; !ok {
			g[p[1]] = unordered_set{
				m:   make(map[int]struct{}),
				set: []int{},
			}
		}
		g[p[0]].Insert(p[1])
		g[p[1]].Insert(p[0])
	}

	var nodes []int
	for _, p := range g {
		nodes = append(nodes, p.First())
	}

	sort.Slice(nodes, func(i, j int) bool {
		return g[i].Size() < g[j].Size()
	})

	for i := 0; i < len(nodes)-1; i++ {
		j := i + 1
		for ; j < len(nodes); j++ {
			if g[nodes[j]].Find(nodes[i]) != g[nodes[j]].End() {
				break
			}
		}
		// 树上除了根节点都应该存在父节点
		if j == len(nodes) {
			return 0
		}
		// nodes[j] 一定是 nodes[i] 的父节点
		t[nodes[j]] = append(t[nodes[j]], nodes[i])
		if g[nodes[j]].Size() == g[nodes[i]].Size() {
			ways = 2
		}
	}

	checkTree(nodes[len(nodes)-1], 0)

	return ways
}
