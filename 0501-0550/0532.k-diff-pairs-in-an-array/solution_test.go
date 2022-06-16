package kdiffpairsinanarray

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, findPairs, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
