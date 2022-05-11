package integerbreak

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithFile(t, integerBreak2, "./in.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
