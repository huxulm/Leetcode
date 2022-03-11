package surroundedregions

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
