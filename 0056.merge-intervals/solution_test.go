package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect [][]int
	}{
		{input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, expect: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{input: [][]int{{1, 4}, {4, 5}}, expect: [][]int{{1, 5}}},
		{input: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, expect: [][]int{{1, 10}}},
	}

	for _, test := range tests {
		if result := merge(test.input); !reflect.DeepEqual(result, test.expect) {
			t.Errorf("Expect %v but got %v", test.expect, result)
		}
	}
}
