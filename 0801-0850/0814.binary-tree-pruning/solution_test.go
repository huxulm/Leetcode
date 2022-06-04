package binarytreepruning

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, pruneTree, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
func TestXxx1(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, pruneTree1, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
