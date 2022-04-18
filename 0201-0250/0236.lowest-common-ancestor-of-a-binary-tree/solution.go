package lowestcommonancestorofabinarytree

import (
	. "lc/structures"
)

func lowestCommonAncestor(root, p, q *TreeNode) (c *TreeNode) {
	if root == nil {
		return root
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	} else {
		return left
	}
}