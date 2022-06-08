package maximumnumberofvowelsinasubstringofgivenlength

func maxVowels(s string, k int) (ans int) {
	var cnts int
	for i := range s {
		ch := s[i]
		if isVowels(ch) {
			cnts++
		}
		if cnts > ans {
			ans = cnts
		}
		if i >= k-1 {
			if isVowels(s[i-k+1]) {
				cnts--
			}
		}
	}
	return
}

func isVowels(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
