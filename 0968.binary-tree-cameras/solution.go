package binarytreecameras

import (
	. "lc/structures"
)

// DFS
// 思路:
// 对于每个节点需要维护三个状态:
//
// 状态 a：root 必须放置摄像头的情况下，覆盖整棵树需要的摄像头数目。
// 状态 b：覆盖整棵树需要的摄像头数目，无论 root 是否放置摄像头。
// 状态 c：覆盖两棵子树需要的摄像头数目，无论节点 root 本身是否被监控到。
func minCameraCover(root *TreeNode) int {
	_, ans, _ := dfs(root)
	return ans
}

var INT_MAX = 1<<31 - 1

func dfs(root *TreeNode) (a, b, c int) {
	if root == nil {
		return INT_MAX / 2, 0, 0
	}
	la, lb, lc := dfs(root.Left)
	ra, rb, rc := dfs(root.Right)

	a = lc + rc + 1
	b = min(a, min(la+rb, ra+lb))
	c = min(a, lb+rb)
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 方法二:
// 首先我们列举一下一个节点可能存在的几种状态：
// 该节点不在监控区域内，称为 uncover 状态；该节点在附近节点的监控范围内，称为 cover 状态；该节点自己装了摄像头，称为 set 状态。
// 如何保证安装的摄像头数量尽可能少呢？显然就是要尽可能分散，让每个摄像头物尽其用。
// 具体来说就是自底向上安装摄像头，在叶子节点的父节点上安装摄像头，然后每隔两层再安装（因为每个摄像头都可以管三层）。
// 那么一个节点在什么情况下需要被安装摄像头呢？显然是当这个节点的子节点处于 uncover 的状态的时候必须安装摄像头，以便覆盖子节点。
// 综上，我们需要利用后序位置自底向上遍历二叉树，同时要利用子节点的状态以及父节点的状态，判断当前节点是否需要安装摄像头。
// 解法中 setCamera 函数就负责按照最优方式给二叉树安装摄像头，同时返回节点的状态。
func minCameraCover1(root *TreeNode) (ans int) {

	// 定义：输入以 root 为根的二叉树，以最优策略在这棵二叉树上放置摄像头，
	// 然后返回 root 节点的情况：
	// 返回 -1 代表 root 为空，返回 0 代表 root 未被 cover，
	// 返回 1 代表 root 已经被 cover，返回 2 代表 root 上放置了摄像头。
	var setCamera func(root *TreeNode, hasParent bool) int
	setCamera = func(root *TreeNode, hasParent bool) int {
		if root == nil {
			return -1
		}
		left, right := setCamera(root.Left, true), setCamera(root.Right, true)
		if left == -1 && right == -1 {
			if hasParent { // 当前节点是叶子节点
				// 有父节点，让父节点cover自己
				return 0
			}
			// 没有父节点，自己set
			ans++
			return 2
		}

		if left == 0 || right == 0 {
			// 左右子树存在没有被 cover 的
			// 必须在当前节点 set 一个摄像头
			ans++
			return 2
		}

		if left == 2 || right == 2 {
			// 左右子树只要有一个 set 了摄像头
			// 当前节点就已经是 cover 状态了
			return 1
		}

		// 剩下 left == 1 && right == 1 的情况
		// 即当前节点的左右子节点都被 cover
		if hasParent {
			// 如果有父节点的话，可以等父节点 cover 自己
			return 0
		} else {
			// 没有父节点，只能自己 set 一个摄像头
			ans++
			return 2
		}
	}

	setCamera(root, false)
	return
}

// 0: no cover
// 1: cover
// 2: camera
var (
	camera int
)

func postorder(cur *TreeNode) int {
	if cur == nil {
		return 1
	}

	left := postorder(cur.Left)
	right := postorder(cur.Right)

	if left == 0 || right == 0 {
		camera++
		return 2
	}

	if left == 2 || right == 2 {
		return 1
	}

	return 0
}

func minCameraCover2(root *TreeNode) int {
	camera = 0
	if root == nil {
		return 0
	}

	left := postorder(root.Left)
	right := postorder(root.Right)
	if left == 0 || right == 0 {
		camera++
	} else if left == 1 && right == 1 {
		camera++
	}

	return camera
}
