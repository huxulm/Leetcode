package removelinkedlistelements

import (
	"reflect"
	"testing"

	. "lc/structures"
	"lc/testutils"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		input, expect *ListNode
		val           int
	}{
		{input: testutils.Arr2List([]int{1, 2, 3, 4, 5, 6}), val: 1, expect: testutils.Arr2List([]int{2, 3, 4, 5, 6})},
		{input: testutils.Arr2List([]int{1, 2, 3, 4, 5, 6}), val: 4, expect: testutils.Arr2List([]int{1, 2, 3, 5, 6})},
		{input: testutils.Arr2List([]int{1, 2, 3, 4, 5, 6}), val: 6, expect: testutils.Arr2List([]int{1, 2, 3, 4, 5})},
		{input: testutils.Arr2List([]int{1, 1, 1, 1}), val: 1, expect: testutils.Arr2List([]int{})},
	}

	for _, test := range tests {
		t.Run("Test removeElements", func(t *testing.T) {
			if result := removeElements(test.input, test.val); !reflect.DeepEqual(testutils.List2Arr(test.expect), testutils.List2Arr(result)) {
				t.Errorf("Input %v val %d expect %v but got %v", testutils.List2Arr(test.input), test.val, testutils.List2Arr(test.expect), testutils.List2Arr(result))
			}
		})
		t.Run("Test removeElements1 (LC official)", func(t *testing.T) {
			if result := removeElements1(test.input, test.val); !reflect.DeepEqual(testutils.List2Arr(test.expect), testutils.List2Arr(result)) {
				t.Errorf("Input %v val %d expect %v but got %v", testutils.List2Arr(test.input), test.val, testutils.List2Arr(test.expect), testutils.List2Arr(result))
			}
		})
	}
}
