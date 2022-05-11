package serializeanddeserializebst

import (
	"fmt"
	. "lc/structures"
	"testing"
)

func TestXxx(t *testing.T) {
	// targetCaseNum := 0
	// if err := testutil.RunLeetCodeClassWithFile(t, Constructor, "./in.txt", targetCaseNum); err != nil {
	// 	t.Fatal(err)
	// }

	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	codec := &Codec{}
	data := codec.serialize(root)
	res := codec.deserialize(data)
	fmt.Println(res)
}
