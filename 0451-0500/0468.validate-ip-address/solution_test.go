package validateipaddress

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func TestValidIPAddress(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, validIPAddress1, "./in.txt", 0); err != nil {
		t.Fatal(err)
	}
}
