package constructstringfrombinarytree

import (
	. "lc/structures"
	"strconv"
)

func tree2str(root *TreeNode) (ans string) {
	traversal(&ans, root)
	return
}

func traversal(r *string, root *TreeNode) {
	if root == nil {
		return
	}

	*r = *r + strconv.Itoa(root.Val)

	if root.Left != nil {
		*r = *r + "("
		traversal(r, root.Left)
		*r = *r + ")"
	} else if root.Left == nil && root.Right != nil {
		*r = *r + "()"
	}

	if root.Right != nil {
		*r = *r + "("
		traversal(r, root.Right)
		*r = *r + ")"
	}
}
