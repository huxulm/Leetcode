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
	n := len(grid)
	return constructHelper(0, 0, n, n, grid)
}

func constructHelper(rowStart, rowEnd, colStart, colEnd int, grid [][]int) *Node {
	if rowEnd == rowStart || colStart == colEnd {
		return nil
	}
	zeroCnt, oneCnt := 0, 0
	for i := rowStart; i < rowEnd; i++ {
		for j := colStart; j < colEnd; j++ {
			if grid[i][j] == 0 {
				zeroCnt++
			} else if grid[i][j] == 1 {
				oneCnt++
			}
			if oneCnt > 0 && zeroCnt > 0 {
				break
			}
		}
	}
	root := &Node{}
	if zeroCnt > 0 && oneCnt > 0 {
		root.IsLeaf = false
		root.TopLeft = constructHelper(rowStart, (rowEnd-rowStart)>>1, colStart, colStart+(colEnd-colStart)>>1, grid)
		root.TopRight = constructHelper(rowStart, (rowEnd-rowStart)>>1, colStart+(colEnd-colStart)>>1, colEnd, grid)
		root.BottomLeft = constructHelper(rowStart+(rowEnd-rowStart)>>1, rowEnd, colStart, colStart+(colEnd-colStart)>>1, grid)
		root.BottomRight = constructHelper(rowStart+(rowEnd-rowStart)>>1, rowEnd, colStart+(colEnd-colStart)>>1, colEnd, grid)
	} else {
		root.IsLeaf = true
		root.Val = zeroCnt != 0
	}
	return root
}
