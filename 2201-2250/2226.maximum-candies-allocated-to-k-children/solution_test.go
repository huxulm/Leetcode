package maximumcandiesallocatedtokchildren

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		candies []int
		k       int
		expect  int
	}{{candies: []int{5, 8, 6}, k: 3, expect: 5},
		{candies: []int{2, 5}, k: 11, expect: 0},
	}

	for _, test := range tests {
		t.Run("#Solution0", func(t *testing.T) {
			if res := maximumCandies(test.candies, int64(test.k)); test.expect != res {
				t.Errorf("Input %v expect %d but got %d", test.candies, test.expect, res)
			}
		})

		t.Run("#Solution1", func(t *testing.T) {
			if res := maximumCandies1(test.candies, int64(test.k)); test.expect != res {
				t.Errorf("Input %v expect %d but got %d", test.candies, test.expect, res)
			}
		})
	}
}
