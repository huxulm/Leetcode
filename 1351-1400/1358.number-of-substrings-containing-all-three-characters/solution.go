package numberofsubstringscontainingallthreecharacters

func numberOfSubstrings(s string) (ans int) {
	n := len(s)
	// a, b, c 个数
	cnts := [3]int{}

	l, r := 0, 0
	for l < n {
		for r < n && (cnts[0] < 1 || cnts[1] < 1 || cnts[2] < 1) {
			if isAbc(s[r]) {
				cnts[s[r]-'a']++
			}
			r++
		}
		if cnts[0] >= 1 && cnts[1] >= 1 && cnts[2] >= 1 {
			ans += n - r + 1
		}
		if isAbc(s[l]) {
			cnts[s[l]-'a']--
		}
		l++
	}
	return
}

func isAbc(ch byte) bool {
	return ch == 'a' || ch == 'b' || ch == 'c'
}
