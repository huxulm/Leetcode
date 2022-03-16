package reversenodesinkgroup

import (
	. "lc/structures"
	"lc/testutils"
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		head   *ListNode
		k      int
		expect *ListNode
	}{
		{
			head:   testutils.Arr2List([]int{1, 2, 3, 4}),
			k:      2,
			expect: testutils.Arr2List([]int{2, 1, 4, 3}),
		},
	}

	for _, test := range tests {
		if result := reverseKGroup(test.head, test.k); !reflect.DeepEqual(testutils.List2Arr(result), testutils.List2Arr(test.expect)) {
			t.Errorf("Input %v", testutils.List2Arr(test.head))
		}
	}
}
