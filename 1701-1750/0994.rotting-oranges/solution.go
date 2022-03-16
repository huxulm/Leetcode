package rottingoranges

import "container/list"

type pair struct {
	x, y int
}

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 广度优先
func orangesRotting(grid [][]int) int {
	// 记录未腐烂和腐烂的数量
	rotted, unrotted := 0, 0
	m, n := len(grid), len(grid[0])

	q := list.New()

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				unrotted++
			} else if grid[i][j] == 2 {
				rotted++
				// 将初始腐烂的橘子放入队列
				q.PushBack(&pair{i, j})
			}
		}
	}

	// 初始没有未腐烂的橘子，直接返回 0
	if unrotted == 0 {
		return 0
	}

	ans := 0
	for q.Len() > 0 {
		// 还有未腐烂的橘子，时间加一
		if unrotted > 0 {
			ans++
		}

		// 层序遍历所有腐烂橘子，搜索各自周围未腐烂的橘子将其腐烂
		for i, sz := 0, q.Len(); i < sz; i++ {
			p := q.Remove(q.Front()).(*pair)
			// 四个方向
			for _, d := range dirs {
				if x, y := p.x+d[0], p.y+d[1]; x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					q.PushBack(&pair{x, y})
					unrotted--
				}
			}
		}
	}

	if unrotted > 0 {
		return -1
	}
	return ans
}
