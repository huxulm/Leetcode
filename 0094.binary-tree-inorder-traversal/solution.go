package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var Traversal func(root *TreeNode, arr *[]int)

	Traversal = func(root *TreeNode, arr *[]int) {
		if root == nil {
			return
		}
		Traversal(root.Left, arr)
		*arr = append(*arr, root.Val)
		Traversal(root.Right, arr)
	}

	arr := make([]int, 0)
	Traversal(root, &arr)
	return arr
}
