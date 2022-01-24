package generateparentheses

func generateParenthesis(n int) []string {
	var res []string
	backtrack(&res, n, 0, 0, 0, make([]byte, 2*n))
	return res
}

func backtrack(res *[]string, n, idx, l, r int, track []byte) {
	if r == n {
		*res = append(*res, string(track))
		return
	}
	// 选择 [idx] 位置的括号
	if l < n && l >= r { // 选左括号
		track[idx] = '('
		backtrack(res, n, idx+1, l+1, r, track)
	}
	if r < n && r < l {
		track[idx] = ')'
		backtrack(res, n, idx+1, l, r+1, track)
	}
}
