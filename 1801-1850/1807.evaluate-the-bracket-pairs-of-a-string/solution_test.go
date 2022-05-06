package evaluatethebracketpairsofastring

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, evaluate, "in.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
