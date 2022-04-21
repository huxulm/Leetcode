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
