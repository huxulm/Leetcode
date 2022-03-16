package simplifypath

import "strings"

func simplifyPath(path string) string {
	names := strings.Split(path, "/")
	stack := []string{}

	for i := range names {
		if names[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if len(names[i]) > 0 && names[i] != "." {
			stack = append(stack, names[i])
		}
	}

	return "/" + strings.Join(stack, "/")
}
