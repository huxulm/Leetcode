package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回溯算法思路
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

// 动态规划思路
func inorderTraversal1(root *TreeNode) []int {
	arr := []int{}
	if root == nil {
		return arr
	}
	arr = append(arr, inorderTraversal1(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, inorderTraversal1(root.Right)...)
	return arr
}
