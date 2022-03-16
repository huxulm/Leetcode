package maximumproductofsplittedbinarytree

import (
	. "lc/structures"
)

func maxProduct(root *TreeNode) (ans int) {
	S := -1
	var postOrder func(root *TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l, r := postOrder(root.Left), postOrder(root.Right)
		if S > 0 {
			ans = max(ans, max((S-l)*l, (S-r)*r))
		}
		return root.Val + l + r
	}
	// 计算 root 大小
	S = postOrder(root)
	// 再次遍历，计算最大的乘积
	postOrder(root)
	ans %= 1000000007
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
