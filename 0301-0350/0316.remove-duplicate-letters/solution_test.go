package removeduplicateletters

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {

	if err := testutil.RunLeetCodeFuncWithFile(t, removeDuplicateLetters, "./in.txt", -1); err != nil {
		t.Fatal(err)
	}
}
