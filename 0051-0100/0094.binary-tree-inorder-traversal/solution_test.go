package binarytreeinordertraversal

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, inorderTraversal2, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
