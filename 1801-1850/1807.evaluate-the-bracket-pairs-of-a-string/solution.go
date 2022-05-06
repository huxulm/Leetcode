package evaluatethebracketpairsofastring

import "strings"

func evaluate(s string, knowledge [][]string) string {
	builder := strings.Builder{}

	dict := map[string]string{}
	for _, p := range knowledge {
		dict[p[0]] = p[1]
	}

	l := -1
	for i := range s {
		if s[i] == '(' && l == -1 {
			l = i
		}
		if l == -1 {
			builder.WriteByte(s[i])
			continue
		}

		if s[i] == ')' {
			if v, ok := dict[s[l+1:i]]; ok {
				builder.WriteString(v)
			} else {
				builder.WriteByte('?')
			}
			l = -1
		}
	}
	return builder.String()
}
