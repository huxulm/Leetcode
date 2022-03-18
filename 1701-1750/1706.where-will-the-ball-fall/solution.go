package wherewilltheballfall

const (
	DOWN, RIGHT, LEFT int = 0, -1, 1
)

// 球只能向右，向左，向下，或被卡住
func findBall(grid [][]int) (ans []int) {
	n, m := len(grid), len(grid[0])
	ans = make([]int, m)

	var dfs func(i, j int, state int) int
	dfs = func(i, j, state int) int {

		// 到边缘被卡住
		if j < 0 || j >= m {
			return -1
		}

		// 顺利从底部出来
		if i == n && j >= 0 && j < m {
			return j
		}

		// 向下/左挡板 => 向右
		if state == DOWN && grid[i][j] == 1 {
			return dfs(i, j+1, RIGHT)
		}
		// 向下/右挡板 => 向左
		if state == DOWN && grid[i][j] == -1 {
			return dfs(i, j-1, LEFT)
		}
		// 向右/左挡板 => 向下 / 向左/右挡板 => 向下
		if (state == RIGHT && grid[i][j] == 1) || state == LEFT && grid[i][j] == -1 {
			return dfs(i+1, j, DOWN)
		}
		// 向左/右挡板  向右/左挡板 被卡住
		return -1
	}

	for j := 0; j < m; j++ {
		ans[j] = dfs(0, j, DOWN)
	}

	return
}

func findBall1(grid [][]int) (ans []int) {
	n, m := len(grid), len(grid[0])
	ans = make([]int, m)

	var getVal func(c int) int

	getVal = func(c int) int {
		r := 0

		for r < n {
			ne := c + grid[r][c]
			if ne < 0 || ne >= m {
				return -1
			}
			// 卡住
			if grid[r][c] != grid[r][ne] {
				return -1
			}
			r, c = r+1, ne
		}

		return c
	}

	for j := 0; j < m; j++ {
		ans[j] = getVal(j)
	}

	return
}
