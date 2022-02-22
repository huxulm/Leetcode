package populatingnextrightpointersineachnode

import (
	"container/list"
	. "lc/structures"
)

// 二叉树层序遍历
func connect(root *Node) *Node {
	q := list.New()
	q.PushBack(root)

	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			prev := q.Remove(q.Front()).(*Node)
			if i < sz-1 {
				next := q.Front().Value.(*Node)
				prev.Next = next
			} else {
				prev.Next = nil
			}
			if prev.Left != nil {
				q.PushBack(prev.Left)
			}
			if prev.Right != nil {
				q.PushBack(prev.Right)
			}
		}
	}
	return root
}
