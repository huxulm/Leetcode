package gridillumination

func gridIllumination1(n int, lamps [][]int, queries [][]int) []int {

	// 开灯 [i,j]: grid[i][]++ grid[][j]++ 对角线++
	// 照亮 grid[i][j] > 0
	// 关灯 [i,j]: grid[i][]-- grid[][j]-- 对角线--
	grid := make([][]int, n)
	for row := range grid {
		grid[row] = make([]int, n)
	}

	state := make([][]bool, n)
	for row := range state {
		state[row] = make([]bool, n)
	}

	// 打开所有lamps
	for i := range lamps {
		openLamp(grid, n, lamps[i][0], lamps[i][1])
		state[lamps[i][0]][lamps[i][1]] = true
	}

	ans := make([]int, len(queries))

	for i := range queries {
		// 查询queries[i]
		ans[i] = queryLamp(grid, queries[i])
		// 关闭queries[i] x9
		closeLamp(grid, state, n, queries[i][0], queries[i][1])

		closeLamp(grid, state, n, queries[i][0]-1, queries[i][1]-1)
		closeLamp(grid, state, n, queries[i][0]-1, queries[i][1])
		closeLamp(grid, state, n, queries[i][0]-1, queries[i][1]+1)

		closeLamp(grid, state, n, queries[i][0], queries[i][1]-1)
		closeLamp(grid, state, n, queries[i][0], queries[i][1]+1)

		closeLamp(grid, state, n, queries[i][0]+1, queries[i][1]-1)
		closeLamp(grid, state, n, queries[i][0]+1, queries[i][1])
		closeLamp(grid, state, n, queries[i][0]+1, queries[i][1]+1)
	}

	return ans
}

func openLamp(grid [][]int, n, i, j int) {
	// 超过边界了
	if i < 0 || j < 0 || i >= n || j >= n {
		return
	}
	// 行
	for col := 0; col < n; col++ {
		grid[i][col]++
	}
	// 列
	for row := 0; row < n; row++ {
		grid[row][j]++
	}
	// 对角线
	// left
	if i <= j {
		for col := j - i; col < n; col++ {
			grid[col+i-j][col]++
		}
	} else {
		for col := 0; col < n-i+j; col++ {
			grid[col+i-j][col]++
		}
	}
	// right
	if i+j < n {
		for col := 0; col < i+j+1; col++ {
			grid[i+j-col][col]++
		}
	} else {
		for col := i + j + 1 - n; col < n; col++ {
			grid[i+j-col][col]++
		}
	}

	grid[i][j] -= 3
}

func closeLamp(grid [][]int, state [][]bool, n, i, j int) {
	// 超过边界了或灯未点亮
	if i < 0 || j < 0 || i >= n || j >= n || !state[i][j] {
		return
	}
	// 行
	for col := 0; col < n; col++ {
		grid[i][col]--
	}
	// 列
	for row := 0; row < n; row++ {
		grid[row][j]--
	}
	// 对角线
	// left
	if i <= j {
		for col := j - i; col < n; col++ {
			grid[col+i-j][col]--
		}
	} else {
		for col := 0; col < n-i+j; col++ {
			grid[col+i-j][col]--
		}
	}
	// right
	if i+j < n {
		for col := 0; col < i+j+1; col++ {
			grid[i+j-col][col]--
		}
	} else {
		for col := i + j + 1 - n; col < n; col++ {
			grid[i+j-col][col]--
		}
	}

	grid[i][j] += 3
	state[i][j] = false
}

func queryLamp(grid [][]int, lamp []int) int {
	if grid[lamp[0]][lamp[1]] > 0 {
		return 1
	}
	return 0
}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	type pair struct{ x, y int }
	points := map[pair]int{}

	row := make(map[int]int)
	col := make(map[int]int)
	diagonal := map[int]int{}
	antiDiagonal := map[int]int{}

	ans := make([]int, len(queries))

	for _, l := range lamps {
		p := pair{l[0], l[1]}
		// 去重
		if _, ok := points[p]; ok {
			continue
		}
		points[p] = 1
		row[p.x]++
		col[p.y]++
		diagonal[p.x-p.y]++
		antiDiagonal[p.x+p.y]++
	}

	for i, q := range queries {
		x, y := q[0], q[1]
		if row[x] > 0 || col[y] > 0 || diagonal[x-y] > 0 || antiDiagonal[x+y] > 0 {
			ans[i] = 1
		}

		// 关灯
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if i < 0 || i >= n || j < 0 || j >= n {
					continue
				}
				p := pair{i, j}
				if _, ok := points[p]; ok {
					row[p.x]--
					col[p.y]--
					diagonal[p.x-p.y]--
					antiDiagonal[p.x+p.y]--
					delete(points, p)
				}
			}
		}
	}

	return ans
}
