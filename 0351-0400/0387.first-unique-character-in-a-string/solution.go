package firstuniquecharacterinastring

func firstUniqChar(s string) int {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}
	for i := range s {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}
