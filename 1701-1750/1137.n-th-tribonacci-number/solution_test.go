package nthtribonaccinumber

import (
	"reflect"
	"testing"
)

func TestTribonacci(t *testing.T) {
	ans := []int{}
	for i := 1; i <= 32; i++ {
		ans = append(ans, tribonacci(i))
	}
	expect := []int{1, 1, 2, 4, 7, 13, 24, 44, 81, 149, 274, 504, 927, 1705, 3136, 5768, 10609, 19513, 35890, 66012, 121415, 223317, 410744, 755476, 1389537, 2555757, 4700770, 8646064, 15902591, 29249425, 53798080, 98950096}
	if !reflect.DeepEqual(ans, expect) {
		t.Errorf("Input n [1,32] expect %v but got %v", expect, ans)
	}
}
