package maximumnumberofachievabletransferrequests

import (
	"fmt"
	"testing"
)

func TestMaximumRequests(t *testing.T) {
	fmt.Println(maximumRequests(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}))
}
