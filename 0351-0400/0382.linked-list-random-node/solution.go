package linkedlistrandomnode

import (
	. "lc/structures"
	"math/rand"
)

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (this *Solution) GetRandom() (ans int) {
	p, i := this.head, 0
	for p != nil {
		i++
		if rand.Intn(i) == 0 {
			ans = p.Val
		}
		p = p.Next
	}
	return
}
