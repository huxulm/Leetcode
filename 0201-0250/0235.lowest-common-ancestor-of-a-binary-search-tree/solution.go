package lowestcommonancestorofabinarysearchtree

import (
	. "lc/structures"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestor1(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if ancestor.Val > p.Val && ancestor.Val > q.Val {
			ancestor = ancestor.Left
		} else if ancestor.Val < p.Val && ancestor.Val < q.Val {
			ancestor = ancestor.Right
		} else {
			return ancestor
		}
	}
}
