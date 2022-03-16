package maxareaofisland

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// 标记访问过的地块
	marked := make([][]bool, m)
	for i := range marked {
		marked[i] = make([]bool, n)
	}

	// 保存每一块土地所在的最大面积
	memo := make([][]int, m)

	for i := range memo {
		memo[i] = make([]int, n)
	}

	ans := -1 << 31

	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var dfs func(area *int, i, j int, memo [][]int, marked [][]bool)
	dfs = func(area *int, i, j int, memo [][]int, marked [][]bool) {
		// 越界或被标记过或不是陆地
		if i < 0 || i >= m || j < 0 || j >= n || marked[i][j] || grid[i][j] == 0 {
			return
		}
		marked[i][j] = true
		// 相邻四个方向有无被搜索过，如果搜索过直接返回值
		for _, d := range dirs {
			if x, y := i+d[0], j+d[1]; x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && memo[x][y] > 0 {
				memo[i][j] = memo[x][y]
				return
			}
		}
		*area++
		// 向四个方向搜索
		for _, d := range dirs {
			dfs(area, i+d[0], j+d[1], memo, marked)
		}
		memo[i][j] = *area
		marked[i][j] = false
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			res := 0
			dfs(&res, i, j, memo, marked)
			ans = max(res, ans)
		}
	}

	return ans
}

func maxAreaOfIsland1(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	ans := 0

	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 越界或被标记过或不是陆地
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		// 标记，确保每个土地被计算一次
		grid[i][j] = 0
		ans := 1
		// 向四个方向搜索
		for _, d := range dirs {
			ans += dfs(i+d[0], j+d[1])
		}
		return ans
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			ans = max(dfs(i, j), ans)
		}
	}

	return ans
}
