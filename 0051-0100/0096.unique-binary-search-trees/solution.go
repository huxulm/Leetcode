package uniquebinarysearchtrees

func numTrees(n int) int {
	// f[i] 表示序列长度为 i 的二叉搜索树方案数
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 以 j 为根, 左子树方案数 * 右子树方案数
			f[i] += f[j-1] * f[i-j]
		}
	}
	return f[n]
}
