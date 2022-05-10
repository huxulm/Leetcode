package catandmouseii

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithFile(t, canMouseWin1, "./in.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
