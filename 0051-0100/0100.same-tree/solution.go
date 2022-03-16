package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	eq := p != nil && q != nil && p.Val == q.Val

	if !eq {
		return false
	}

	return eq && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
