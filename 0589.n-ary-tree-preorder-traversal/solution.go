package narytreepreordertraversal

//  Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 方法一: DFS/递归
func preorder(root *Node) (ans []int) {
	traversal(&ans, root)
	return
}

func traversal(res *[]int, root *Node) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, s := range root.Children {
		traversal(res, s)
	}
}
