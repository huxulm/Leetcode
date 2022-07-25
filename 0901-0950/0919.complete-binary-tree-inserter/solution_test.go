package completebinarytreeinserter

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestCBTInserter(t *testing.T) {
	if err := testutil.RunLeetCodeClassWithFile(t, Constructor, "./in.txt", -1); err != nil {
		t.Errorf("error!")
	}
}
