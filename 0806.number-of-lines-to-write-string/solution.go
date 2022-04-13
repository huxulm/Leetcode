package numberoflinestowritestring

func numberOfLines(widths []int, s string) (ans []int) {
	ans = make([]int, 2)

	width := 100
	ans[0] = 1
	for i := range s {
		w := widths[s[i]-'a']
		if width < w {
			width = 100
			ans[0]++
		}
		width -= w
	}
	ans[1] = 100 - width
	return
}
