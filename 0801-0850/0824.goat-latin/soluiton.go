package goatlatin

import "strings"

func toGoatLatin(sentence string) string {
	ans := &strings.Builder{}
	sp := strings.Split(sentence, " ")
	for i := range sp {
		if strings.IndexByte("aeiou", byte(sp[i][0]|' ')) != -1 {
			ans.WriteString(sp[i])
		} else {
			ans.WriteString(sp[i][1:])
			ans.WriteString(sp[i][:1])
		}
		ans.WriteString("ma")

		for j := 0; j <= i; j++ {
			ans.WriteString("a")
		}
		if i != len(sp)-1 {
			ans.WriteString(" ")
		}
	}
	return ans.String()
}

func toGoatLatin1(s string) string {
	n := len(s)
	last, cnt := 0, 0
	ans := &strings.Builder{}
	for i := 0; i < n; i++ {

		for i < n && s[i] != ' ' {
			i++
		}
		cnt++
		w := s[last:i]
		last = i + 1
		if strings.IndexByte("aeiou", byte(w[0]|' ')) != -1 {
			ans.WriteString(w)
		} else {
			ans.WriteString(w[1:])
			ans.WriteString(w[:1])
		}
		ans.WriteString("ma")

		for j := 0; j <= cnt; j++ {
			ans.WriteString("a")
		}
		if i != n {
			ans.WriteString(" ")
		}
		last = i - 1
	}
	return ans.String()
}
