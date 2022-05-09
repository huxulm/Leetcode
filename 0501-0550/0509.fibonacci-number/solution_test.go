package fibonaccinumber

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, fib1, "./in.txt", -1); err != nil {
		t.Fatal(err)
	}
}
