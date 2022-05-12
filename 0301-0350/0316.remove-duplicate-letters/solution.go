package removeduplicateletters

func removeDuplicateLetters(s string) string {
	cnts := [26]int{}
	inQ := [26]bool{}

	for i := range s {
		cnts[s[i]-'a']++
	}

	q := []byte{}
	for i := range s {
		sz := len(q)
		ch := byte(s[i])
		cnts[ch-'a']--
		if sz == 0 {
			q = append(q, ch)
			inQ[ch-'a'] = true
			continue
		}
		for ; len(q) > 0 && !inQ[ch-'a'] && q[sz-1] > ch && cnts[q[sz-1]-'a'] > 0; sz = len(q) {
			inQ[q[sz-1]-'a'] = false
			q = q[:sz-1]
		}
		if !inQ[ch-'a'] {
			q = append(q, ch)
			inQ[ch-'a'] = true
		}
	}
	return string(q)
}
