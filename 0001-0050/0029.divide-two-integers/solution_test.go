package dividetwointegers

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, divide1, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
