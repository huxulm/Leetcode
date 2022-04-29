package constructquadtree

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return dfs(grid, 0, 0, len(grid))
}

func dfs(g [][]int, r, c, n int) *Node {
	if n == 0 {
		return nil
	}
	zero, one := isAllsame(g, r, c, n)
	if zero {
		return &Node{Val: false, IsLeaf: true}
	}
	if one {
		return &Node{Val: true, IsLeaf: true}
	}
	return &Node{
		true,
		false,
		dfs(g, r, c, n/2),
		dfs(g, r, c+n/2, n/2),
		dfs(g, r+n/2, c, n/2),
		dfs(g, r+n/2, c+n/2, n/2),
	}
}

// 判断 n x n 是否全 1 或 0
func isAllsame(g [][]int, r, c, n int) (zero, one bool) {
	sum := 0
	for i := r; i < r+n; i++ {
		for j := c; j < c+n; j++ {
			sum += g[i][j]
		}
	}
	return sum == 0, sum == n*n
}

// 方法二: 前缀和优化

func construct1(grid [][]int) *Node {

	N := len(grid)
	pre := make([][]int, N+1)
	pre[0] = make([]int, N+1)
	for i, row := range grid {
		pre[i+1] = make([]int, N+1)
		for j, v := range row {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + v
		}
	}

	var dfs func(g [][]int, r, c, n int) *Node
	dfs = func(g [][]int, r, c, n int) *Node {
		if n == 0 {
			return nil
		}
		total := pre[r+n][c+n] - pre[r+n][c] - pre[r][c+n] + pre[r][c]
		zero, one := total == 0, total == n*n
		if zero {
			return &Node{Val: false, IsLeaf: true}
		}
		if one {
			return &Node{Val: true, IsLeaf: true}
		}
		return &Node{
			true,
			false,
			dfs(g, r, c, n/2),
			dfs(g, r, c+n/2, n/2),
			dfs(g, r+n/2, c, n/2),
			dfs(g, r+n/2, c+n/2, n/2),
		}
	}
	return dfs(grid, 0, 0, len(grid))
}
