package matchstickstosquare

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestXxx(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, makesquare, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
func TestXxx1(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, makesquare1, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
