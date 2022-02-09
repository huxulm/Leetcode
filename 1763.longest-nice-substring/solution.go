package longestnicesubstring

import "unicode"

// 最长美好子串
func longestNiceSubstring(s string) string {
	// aAbB
	ans := ""

	n := len(s)

	for i := 0; i < n; i++ {
		lower, upper := 0, 0
		for j := i; j < n; j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}
			if lower == upper && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}
