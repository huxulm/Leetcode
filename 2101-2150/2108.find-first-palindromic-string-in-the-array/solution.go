package findfirstpalindromicstringinthearray

func firstPalindrome(words []string) (ans string) {
next:
	for _, w := range words {
		n := len(w)
		if n == 1 {
			return w
		} else {
			for i := 0; i < n>>1; i++ {
				if w[i] != w[n-i-1] {
					continue next
				}
			}
			return w
		}
	}
	return
}
