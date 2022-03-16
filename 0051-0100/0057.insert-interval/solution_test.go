package insertinterval

import (
	"fmt"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	intervals := [][]int{{1, 3}, {4, 5}, {7, 9}}
	newInterval := []int{7, 9}
	fmt.Println(insert(intervals, newInterval))
}
