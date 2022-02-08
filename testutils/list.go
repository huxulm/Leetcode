package testutils

import (
	. "lc/structures"
)

func Arr2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{}
	var prev = head
	for i, v := range arr {
		if i == 0 {
			prev.Val = v
			continue
		}
		prev.Next = &ListNode{Val: v}
		prev = prev.Next
	}
	return head
}

func List2Arr(list *ListNode) []int {
	arr := make([]int, 0)
	for v := list; v != nil; v = v.Next {
		arr = append(arr, v.Val)
	}
	return arr
}
