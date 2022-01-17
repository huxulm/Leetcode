package convertsortedarraytobinarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}
	if n == 2 {
		return &TreeNode{Val: nums[1], Left: &TreeNode{Val: nums[0]}}
	}
	// n >= 3
	mid := n >> 1
	root := &TreeNode{Val: nums[mid]}
	// 构建左侧平衡树
	root.Left = sortedArrayToBST(nums[:mid])
	// 构建右侧平衡树
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// LC-official题解
func sortedArrayToBST1(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 总是选择中间位置左边的数字作为根节点
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
