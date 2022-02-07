package removeduplicatesfromsortedlistii

import (
	"reflect"
	"testing"
)

func arrToLinkedList(arr []int) *ListNode {
	n := len(arr)
	if n < 1 {
		return nil
	}
	head := new(ListNode)
	head.Val = arr[0]

	prev := head

	for i := 1; i < n; i++ {
		next := &ListNode{Val: arr[i]}
		prev.Next = next
		prev = next
	}

	return head
}

func list2Arr(head *ListNode) []int {

	if head == nil {
		return nil
	}

	arr := []int{}

	for n := head; n != nil; n = n.Next {
		arr = append(arr, n.Val)
	}

	return arr
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		head   []int
		expect []int
	}{
		{head: []int{1, 1, 2, 3, 3, 4, 4, 5}, expect: []int{2, 5}},
		{head: []int{1, 2, 3, 3}, expect: []int{1, 2}},
		{head: []int{1, 1, 1, 3, 3}, expect: nil},
	}

	for _, test := range tests {
		if result := deleteDuplicates(arrToLinkedList(test.head)); !reflect.DeepEqual(list2Arr(result), test.expect) {
			t.Errorf("Input %v expect %v got %v", test.head, test.expect, list2Arr(result))
		}
	}
}
