package verifyinganaliendictionary

func isAlienSorted(words []string, order string) bool {
	dict := [26]byte{}
	for i, ch := range order {
		dict[ch-'a'] = byte(i + 1)
	}
	var comp func(s1, s2 string) bool
	comp = func(s1, s2 string) bool {
		m, n := len(s1), len(s2)
		i := 0
		for ; i < m && i < n; i++ {
			if dict[s1[i]-'a'] == dict[s2[i]-'a'] {
				continue
			} else if dict[s1[i]-'a'] < dict[s2[i]-'a'] {
				return true
			} else {
				return false
			}
		}
		if m <= n {
			return true
		} else {
			return false
		}
	}

	n := len(words)
	for i := range words[:n-1] {
		if comp(words[i], words[i+1]) {
			continue
		} else {
			return false
		}
	}
	return true
}
