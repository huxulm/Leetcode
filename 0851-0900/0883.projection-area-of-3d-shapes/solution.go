package projectionareaof3dshapes

func projectionArea(grid [][]int) (ans int) {
	n := len(grid)
	for i := 0; i < n; i++ {
		max_x := grid[i][0]
		max_y := grid[0][i]
		for j := 0; j < n; j++ {
			if grid[i][j] > max_x {
				max_x = grid[i][j]
			}
			if grid[j][i] > max_y {
				max_y = grid[j][i]
			}
			if grid[i][j] > 0 {
				ans++
			}
		}
		ans += max_x + max_y
	}
	return
}
