package stepbystepdirectionsfromabinarytreenodetoanother

import (
	. "lc/structures"
)

// 方法一: 最近公共祖先
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var startPath, destPath []byte
	path := []byte{}

	var traversal func(path *[]byte, root *TreeNode)
	traversal = func(path *[]byte, root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == startValue {
			startPath = make([]byte, len(*path))
			copy(startPath, *path)
		} else if root.Val == destValue {
			destPath = make([]byte, len(*path))
			copy(destPath, *path)
		}

		*path = append(*path, 'L')
		traversal(path, root.Left)
		*path = (*path)[:len(*path)-1]

		*path = append(*path, 'R')
		traversal(path, root.Right)
		*path = (*path)[:len(*path)-1]
	}

	traversal(&path, root)

	idx, n1, n2 := 0, len(startPath), len(destPath)
	for idx < n1 && idx < n2 && startPath[idx] == destPath[idx] {
		idx++
	}
	startPath, destPath = startPath[idx:], destPath[idx:]
	for i := range startPath {
		startPath[i] = 'U'
	}
	return string(startPath) + string(destPath)
}
