package surroundedregions

import "container/list"

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#'
		for _, d := range dirs {
			dfs(i+d[0], j+d[1])
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// 方法二: 广度优先搜素
type pair struct{ x, y int }

func solve1(board [][]byte) {
	q := list.New()
	m, n := len(board), len(board[0])

	// 0, n-1列
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			q.PushBack(&pair{i, 0})
		}
		if board[i][n-1] == 'O' {
			q.PushBack(&pair{i, n - 1})
		}
	}
	// 0, m-1行
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			q.PushBack(&pair{0, j})
		}
		if board[m-1][j] == 'O' {
			q.PushBack(&pair{m - 1, j})
		}
	}

	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			p := q.Remove(q.Front()).(*pair)
			board[p.x][p.y] = '#'
			for _, d := range dirs {
				if x, y := p.x+d[0], p.y+d[1]; x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'O' {
					q.PushBack(&pair{x, y})
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
