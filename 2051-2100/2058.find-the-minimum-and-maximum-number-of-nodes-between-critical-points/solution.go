package findtheminimumandmaximumnumberofnodesbetweencriticalpoints

import (
	. "lc/structures"
)

func nodesBetweenCriticalPoints(head *ListNode) []int {
	minDist, maxDist := -1, -1
	first, last, pos := -1, -1, 0

	cur := head
	for cur.Next != nil && cur.Next.Next != nil {
		x, y, z := cur.Val, cur.Next.Val, cur.Next.Next.Val
		if (x < y && z < y) || (x > y && z > y) {
			if first == -1 {
				first, last = pos, pos
			} else {
				if minDist == -1 {
					minDist, maxDist = pos-first, pos-first
				} else {
					minDist, maxDist = Min(minDist, pos-last), pos-first
				}
				last = pos
			}
		}
		pos++
		cur = cur.Next
	}
	return []int{minDist, maxDist}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
