package rotatestring

import "strings"

func rotateString(s string, goal string) bool {
	m, n := len(s), len(goal)

	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		if s[i:] == goal[:m-i] && s[:i] == goal[m-i:] {
			return true
		}
	}
	return false
}

func rotateString1(s, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
