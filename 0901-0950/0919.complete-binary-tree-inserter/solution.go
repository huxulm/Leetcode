package completebinarytreeinserter

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

type TreeNode = testutil.TreeNode

type CBTInserter struct {
	root  *TreeNode
	cur   []*TreeNode
	next  []*TreeNode
	index int
}

func Constructor(root *TreeNode) CBTInserter {
	t := CBTInserter{cur: []*TreeNode{root}, root: root}
	t.move()
	return t
}

// move to last level
func (t *CBTInserter) move() {
	for t.index < len(t.cur) {
		cur := t.cur[t.index]
		if cur.Left != nil {
			t.next = append(t.next, cur.Left)
		} else {
			break
		}
		if cur.Right != nil {
			t.next = append(t.next, cur.Right)
			t.index++
		} else {
			break
		}
		if t.index == len(t.cur) {
			t.cur, t.next, t.index = t.next, nil, 0
		}
	}
}

func (t *CBTInserter) Insert(val int) int {
	n := TreeNode{Val: val}
	cur := t.cur[t.index]
	if cur.Left == nil {
		cur.Left = &n
	} else {
		cur.Right = &n
		t.index++
	}
	t.next = append(t.next, &n)
	if t.index == len(t.cur) {
		t.cur, t.next, t.index = t.next, nil, 0
	}
	return cur.Val
}

func (t *CBTInserter) Get_root() *TreeNode {
	return t.root
}
