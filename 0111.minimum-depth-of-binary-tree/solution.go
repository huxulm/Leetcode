package minimumdepthofbinarytree

import "container/list"

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
	q := list.New()
	q.PushBack(root)

	minDepth := 1

	for q.Len() > 0 {
		sz := q.Len()
		minDepth++
		for i := 0; i < sz; i++ {
			root = q.Remove(q.Front()).(*TreeNode)
			if root.Left == nil && root.Right == nil {
				goto min
			}
			if root.Left != nil {
				q.PushBack(root.Left)
			}
			if root.Right != nil {
				q.PushBack(root.Right)
			}
		}
	}

min:
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
