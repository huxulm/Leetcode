package matrix

import "container/list"

// 左、右、下、上四个方向
var dirs = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

type pair struct {
	x, y int
}

// 方法一: 广度优先搜索(BFS)
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	var vis = make([][]bool, m)
	var dist = make([][]int, m)

	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
		dist[i] = make([]int, n)
	}

	// 队列
	q := list.New()

	// 将所有的 0 添加进初始队列中
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q.PushBack(&pair{i, j})
				vis[i][j] = true
			}
		}
	}

	// 广度优先搜索
	for q.Len() > 0 {
		// 出队
		p := q.Remove(q.Front()).(*pair)
		for _, d := range dirs {
			if x, y := p.x+d[0], p.y+d[1]; x >= 0 && x < m && y >= 0 && y < n && !vis[x][y] {
				dist[x][y] = dist[p.x][p.y] + 1
				q.PushBack(&pair{x, y})
				vis[x][y] = true
			}
		}
	}

	return dist
}

// 方法二: 动态规划
//
// 思路:
// 从一个固定的 1 走到任意一个 0，在距离最短的前提下可能有四种方法：
// 		1. 只有 水平向左移动 和 竖直向上移动；
// 		2. 只有 水平向左移动 和 竖直向下移动；
// 		3. 只有 水平向右移动 和 竖直向上移动；
// 		4. 只有 水平向右移动 和 竖直向下移动。

var INT_MAX = 1<<31 - 1

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func updateMatrix1(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dist := make([][]int, m)

	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			// 初值设为很大
			dist[i][j] = INT_MAX >> 1

			// 如果(i, j)处值为 0, 距离为0
			if mat[i][j] == 0 {
				dist[i][j] = 0
			}
		}
	}

	// 只有 水平向左移动 和 竖直向上移动，注意动态规划的计算顺序
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j-1 >= 0 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}
	// 只有 水平向左移动 和 竖直向下移动，注意动态规划的计算顺序
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i+1 < m {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j-1 >= 0 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}
	// 只有 水平向右移动 和 竖直向上移动，注意动态规划的计算顺序
	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if i-1 >= 0 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j+1 < n {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}
	// 只有 水平向右移动 和 竖直向下移动，注意动态规划的计算顺序
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j+1 < n {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}

	return dist
}
