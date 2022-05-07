package minimumdepthofbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS+队列的层序遍历
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	minDepth := 0
	for ; q != nil; minDepth++ {
		tmp := q
		q = nil
		for _, root := range tmp {
			if root.Left == nil && root.Right == nil {
				return minDepth + 1
			}
			if root.Left != nil {
				q = append(q, root.Left)
			}
			if root.Right != nil {
				q = append(q, root.Right)
			}
		}
	}
	return minDepth
}

// DFS
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := 1<<31 - 1
	if root.Left != nil {
		minD = min(minDepth1(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth1(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
