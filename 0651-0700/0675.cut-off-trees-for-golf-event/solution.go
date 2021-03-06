package cutofftreesforgolfevent

import "sort"

type pair struct{ dis, x, y int }

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func cutOffTree(forest [][]int) (ans int) {

	trees := []pair{}

	for i := range forest {
		for j, h := range forest[i] {
			if h > 1 {
				trees = append(trees, pair{h, i, j})
			}
		}
	}

	sort.Slice(trees, func(i, j int) bool { return trees[i].dis < trees[j].dis })

	m, n := len(forest), len(forest[0])

	var bfs func(sx, sy, tx, ty int) int

	bfs = func(sx, sy, tx, ty int) (d int) {
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []pair{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, d := range dirs {
				if x, y := p.x+d.x, p.y+d.y; x >= 0 && x < m && y >= 0 && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, pair{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}

	preX, preY := 0, 0
	for _, t := range trees {
		d := bfs(preX, preY, t.x, t.y)
		if d < 0 {
			return -1
		}
		ans += d
		preX, preY = t.x, t.y
	}
	return
}
