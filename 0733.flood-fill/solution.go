package floodfill

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m, n := len(image), len(image[0])
	marked := make([][]bool, m)
	for i := range marked {
		marked[i] = make([]bool, n)
	}

	ans := make([][]int, m)
	copy(ans, image)

	var dfs func(i, j int, color int)
	dfs = func(i, j, color int) {
		if i < 0 || i >= m || j < 0 || j >= n || marked[i][j] || image[i][j] != color {
			return
		}
		marked[i][j] = true
		ans[i][j] = newColor
		for _, d := range dirs {
			dfs(i+d[0], j+d[1], color)
		}
		marked[i][j] = false
	}
	dfs(sr, sc, image[sr][sc])
	return ans
}
