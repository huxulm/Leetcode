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
