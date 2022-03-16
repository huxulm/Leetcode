package romantointeger

func romanToInt(s string) int {

	var m = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	var ans int

	l := len(s)
	for i := 0; i < l; i++ {
		if i < l-1 && m[s[i]] < m[s[i+1]] {
			ans -= m[s[i]]
		} else {
			ans += m[s[i]]
		}
	}
	return ans
}
