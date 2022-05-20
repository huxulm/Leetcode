package findtheduplicatenumber

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	t.Run("#1", func(t *testing.T) {
		if err := testutil.RunLeetCodeFuncWithFile(t, findDuplicate, "./in.txt", 0); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("#2", func(t *testing.T) {
		if err := testutil.RunLeetCodeFuncWithFile(t, findDuplicate1, "./in.txt", 0); err != nil {
			t.Fatal(err)
		}
	})
}
