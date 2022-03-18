package maximumnumberofballoons

func maxNumberOfBalloons(text string) int {
	// balloon
	cnt := make([]int, 5)

	for i := range text {
		if text[i] == 'b' {
			cnt[0]++
		}
		if text[i] == 'a' {
			cnt[1]++
		}
		if text[i] == 'l' {
			cnt[2]++
		}
		if text[i] == 'o' {
			cnt[3]++
		}
		if text[i] == 'n' {
			cnt[4]++
		}
	}

	return min(cnt[0], min(min(cnt[1], min(cnt[2], cnt[3])/2), cnt[4]))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
