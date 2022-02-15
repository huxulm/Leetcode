package validanagram

func isAnagram(s string, t string) bool {
	var m1, m2 [26]int
	for i := range s {
		m1[s[i]-'a']++
	}
	for i := range t {
		m2[t[i]-'a']++
	}
	for i := range m1 {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}
