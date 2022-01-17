package longestcommonprefix

// func longestCommonPrefix(strs []string) string {
// 	// LCP <= minLen(s)
// 	var minLen = 1<<31 - 1
// 	for i := range strs {
// 		if l := len(strs[i]); l < minLen {
// 			minLen = l
// 		}
// 	}

// 	var lcp int
// 	for i := 0; i < minLen; i++ {
// 		var prefix = strs[0][i]
// 		var valid = true
// 		for j := range strs {
// 			if strs[j][i] != prefix {
// 				valid = false
// 				break
// 			}
// 		}
// 		if valid {
// 			lcp = i
// 		} else {
// 			break
// 		}
// 	}
// 	return strs[0][0:lcp]
// }

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := range strs {
		prefix = lcp(strs[i], prefix)
	}
	return prefix
}

func lcp(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	res := []byte{}

	for i, j := 0, 0; i < l1 && j < l2; i, j = i+1, j+1 {
		if s1[i] != s2[j] {
			break
		}
		res = append(res, s1[i])
	}
	return string(res)
}
