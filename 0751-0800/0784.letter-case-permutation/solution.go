package lettercasepermutation

func letterCasePermutation(s string) (ans []string) {
	n := len(s)
	track := make([]byte, n)
	var dfs func(i int, track []byte)
	dfs = func(i int, track []byte) {
		if i == n {
			ans = append(ans, string(track))
			return
		}
		if s[i] <= '9' && s[i] >= '0' {
			track[i] = s[i]
			dfs(i+1, track)
		} else {
			// 考虑转或不转
			// 不转
			track[i] = s[i]
			dfs(i+1, track)

			if s[i] >= 'a' && s[i] <= 'z' { // 转大写
				track[i] = s[i] - 'a' + 'A'
			} else { // 转小写
				track[i] = s[i] - 'A' + 'a'
			}

			dfs(i+1, track)
		}
	}
	dfs(0, track)
	return
}
