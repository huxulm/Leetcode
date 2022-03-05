package backspacestringcompare

func backspaceCompare(s string, t string) bool {
	ls, lt := len(s), len(t)
	sta1, sta2 := make([]byte, ls), make([]byte, lt)

	p1, p2 := 0, 0

	for i := range s {
		if s[i] != '#' {
			sta1[p1] = s[i]
			p1++
		} else { // '#'
			if p1 > 0 {
				p1--
			}
		}
	}
	for i := range t {
		if t[i] != '#' {
			sta2[p2] = t[i]
			p2++
		} else { // '#'
			if p2 > 0 {
				p2--
			}
		}
	}
	return string(sta1[:p1]) == string(sta2[:p2])
}

// 方法二: 双指针从后往前遍历
func backspaceCompare1(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
