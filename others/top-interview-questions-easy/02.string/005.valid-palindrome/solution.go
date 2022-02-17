package validpalindrome

import "unicode"

func isValidChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}

func isPalindrome(s string) bool {
	n := len(s)
	l, r := 0, n-1

	for l < r {
		if !isValidChar(s[l]) {
			l++
			continue
		}
		if !isValidChar(s[r]) {
			r--
			continue
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}
