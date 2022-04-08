package narytreelevelordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (ans [][]int) {

	if root == nil {
		return
	}

	q := []*Node{root}

	idx := 0
	for len(q) > 0 {
		tmp := q
		q = nil

		ans = append(ans, []int{})

		for _, n := range tmp {
			ans[idx] = append(ans[idx], n.Val)
			for _, c := range n.Children {
				if c != nil {
					q = append(q, c)
				}
			}
		}
		idx++
	}
	return
}
