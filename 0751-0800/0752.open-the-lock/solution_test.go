package openthelock

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, openLock, "./in.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
	if err := testutil.RunLeetCodeFuncWithFile(t, openLock1, "./in.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
