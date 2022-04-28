package pacificatlanticwaterflow

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	tests := []struct {
		heights [][]int
		expect  [][]int
	}{
		{
			heights: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			expect:  [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
	}

	for _, test := range tests {
		if res := pacificAtlantic(test.heights); !reflect.DeepEqual(res, test.expect) {
			t.Errorf("Input %v expect %v but got %v", test.heights, test.expect, res)
		}
	}
}
