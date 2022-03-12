package narytreepostordertraversal

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (ans []int) {
	traversal(&ans, root)
	return
}

func traversal(res *[]int, root *Node) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		if child != nil {
			traversal(res, child)
		}
	}
	*res = append(*res, root.Val)
}
