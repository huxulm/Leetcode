package maximizetheconfusionofanexam

// 方法一: 滑动窗口
func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveAnswersChar(answerKey, k, 'T'), maxConsecutiveAnswersChar(answerKey, k, 'F'))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 计算窗口内 ch 个数 sum <= k 时窗口右侧扩大，sum > k 时 缩小左侧
func maxConsecutiveAnswersChar(s string, k int, ch byte) (ans int) {
	sum, n := 0, len(s)
	l, r := 0, 0
	for r < n {
		if s[r] == ch {
			sum++
		}

		for l < r && sum > k {
			if s[l] == ch {
				sum--
			}
			l++
		}

		if r-l+1 > ans {
			ans = r - l + 1
		}

		r++
	}

	return
}
