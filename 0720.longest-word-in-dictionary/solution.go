package longestwordindictionary

import "sort"

// 示例1
// 输入：words = ["w","wo","wor","worl", "world"]
// 输出："world"
// 解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。

// 示例2
// words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
// 输出："apple"
// 解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"
func longestWord(words []string) (ans string) {

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	m := map[string]bool{}
	for i := range words {
		m[words[i]] = true
	}

	for i := range words {
		if len(words[i]) > len(ans) || (len(words[i]) == len(ans) && words[i] < ans) {
			good := true
			// 检查当前单词的每一个前缀是否都存在
			for j := 1; j < len(words[i]); j++ {
				if !m[words[i][:j]] {
					good = false
				}
				break
			}
			if good {
				ans = words[i]
			}
		}
	}

	return
}
