package reverselinkedlist

import (
	"reflect"
	"testing"

	. "lc/structures"
	"lc/testutils"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{input: testutils.Arr2List([]int{1, 2, 3, 4}), expect: testutils.Arr2List([]int{4, 3, 2, 1})},
		{input: testutils.Arr2List([]int{4, 3, 2, 1}), expect: testutils.Arr2List([]int{1, 2, 3, 4})},
		{input: testutils.Arr2List([]int{}), expect: testutils.Arr2List(nil)},
	}

	for _, test := range tests {
		if result := reverseList(test.input); !reflect.DeepEqual(testutils.List2Arr(result), testutils.List2Arr(test.expect)) {
			t.Errorf("Input %v expect %v but got %v", testutils.List2Arr(test.input), testutils.List2Arr(test.expect), testutils.List2Arr(result))
		}
	}
}
