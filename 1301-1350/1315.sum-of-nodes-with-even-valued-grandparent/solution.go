package sumofnodeswithevenvaluedgrandparent

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func sumEvenGrandparent(root *TreeNode) (ans int) {
	q := []*TreeNode{root}
	m := map[*TreeNode]*TreeNode{}
	for depth := 0; q != nil; depth++ {
		tmp := q
		q = nil
		for _, n := range tmp {
			if n.Left != nil {
				if pp, ok := m[n]; ok && pp.Val%2 == 0 {
					ans += n.Left.Val
				}
				q = append(q, n.Left)
				m[n.Left] = n
			}
			if n.Right != nil {
				if pp, ok := m[n]; ok && pp.Val%2 == 0 {
					ans += n.Right.Val
				}
				q = append(q, n.Right)
				m[n.Right] = n
			}
		}
	}
	return
}
