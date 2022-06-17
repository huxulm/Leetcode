package findandreplacepattern

func findAndReplacePattern(words []string, pattern string) (ans []string) {
	for _, w := range words {
		if len(w) != len(pattern) {
			continue
		}
		m1, m2 := map[byte]byte{}, map[byte]byte{}
		var valid bool = true
		for i := 0; i < len(w); i++ {
			a, b := w[i], pattern[i]
			if v, ok := m1[a]; ok && v != byte(b) {
				valid = false
				break
			}
			if v, ok := m2[b]; ok && v != byte(a) {
				valid = false
				break
			}
			// a => b, b => a
			m1[a], m2[b] = b, a
		}

		if valid {
			ans = append(ans, w)
		}
	}
	return
}

func match(pat, str string) bool {
	if len(pat) != len(str) {
		return false
	}
	m := map[byte]byte{}
	for i := range str {
		if v, ok := m[str[i]]; !ok {
			m[str[i]] = pat[i]
		} else if v != pat[i] {
			return false
		}
	}
	return true
}

func findAndReplacePattern1(words []string, pattern string) (ans []string) {
	for i := range words {
		if match(pattern, words[i]) && match(words[i], pattern) {
			ans = append(ans, words[i])
		}
	}
	return
}
