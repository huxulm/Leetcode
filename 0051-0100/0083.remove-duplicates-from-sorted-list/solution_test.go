package removeduplicatesfromsortedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{input: []int{1, 1, 2, 2, 3, 3}, expect: []int{1, 2, 3}},
	}

	for _, c := range tests {
		if r := deleteDuplicates(arr2List(c.input)); !reflect.DeepEqual(list2Arr(r), c.expect) {
			t.Errorf("Input: %v, expect: %v but got: %v", c.input, c.expect, list2Arr(r))
		} else {
			fmt.Printf("input: %v result: %v", c.input, list2Arr(r))
		}
	}
}

func arr2List(arr []int) *ListNode {
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

func list2Arr(list *ListNode) []int {
	arr := make([]int, 0)
	for v := list; v != nil; v = v.Next {
		arr = append(arr, v.Val)
	}
	return arr
}
