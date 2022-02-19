package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n2 < n1 {
		return false
	}

	cnt_s1, cnt_s2 := make([]int, 26), make([]int, 26)

	for i := range s1 {
		cnt_s1[s1[i]-'a']++
	}

	var checkEqual = func(a, b []int) bool {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < n2; i++ {
		if i < n1 {
			cnt_s2[s2[i]-'a']++
			if i == n1-1 && checkEqual(cnt_s1, cnt_s2) {
				return true
			}
		} else {
			cnt_s2[s2[i]-'a']++
			cnt_s2[s2[i-n1]-'a']--
			if checkEqual(cnt_s1, cnt_s2) {
				return true
			}
		}
	}

	return false
}
