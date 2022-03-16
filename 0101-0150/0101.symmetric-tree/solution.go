package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}
		// 两个根节点需要相同
		if left.Val != right.Val {
			return false
		}
		// 左右子节点需要对称相同
		return check(left.Right, right.Left) && check(left.Left, right.Right)
	}

	return check(root.Left, root.Right)
}
