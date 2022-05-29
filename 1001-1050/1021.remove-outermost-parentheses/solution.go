package removeoutermostparentheses

func removeOuterParentheses(s string) (ans string) {
	l, r := 0, 0
	for i := range s {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			ans += s[i-l-r+2 : i]
			l, r = 0, 0
		}
	}
	return
}
