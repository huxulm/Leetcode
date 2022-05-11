package serializeanddeserializebst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	. "lc/structures"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 1[9,8],2[3[null,5],null]
	ans := strings.Builder{}

	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans.WriteString(strconv.Itoa(root.Val))
		if root.Left == root.Right {
			return
		}
		if root.Left != nil {
			ans.WriteByte(' ')
			preOrder(root.Left)
		}
		if root.Right != nil {
			ans.WriteByte(' ')
			preOrder(root.Right)
		}
	}
	preOrder(root)
	return ans.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	sp := strings.Split(data, " ")
	rootVal, _ := strconv.Atoi(sp[0])
	root := &TreeNode{Val: rootVal}
	i := 1
	for ; i < len(sp); i++ {
		if val, _ := strconv.Atoi(sp[i]); val >= rootVal {
			break
		}
	}
	root.Left = this.deserialize(strings.Join(sp[1:i], " "))
	root.Right = this.deserialize(strings.Join(sp[i:], " "))
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
