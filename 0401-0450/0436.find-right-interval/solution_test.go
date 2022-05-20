package findrightinterval

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, findRightInterval, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
