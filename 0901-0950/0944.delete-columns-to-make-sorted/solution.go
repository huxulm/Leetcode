package deletecolumnstomakesorted

func minDeletionSize(strs []string) (ans int) {
	m, n := len(strs), len(strs[0])
	for j := 0; j < n; j++ {
		for i := 0; i < m-1; i++ {
			if strs[i][j] > strs[i+1][j] {
				ans++
				break
			}
		}
	}
	return
}
