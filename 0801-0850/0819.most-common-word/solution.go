package mostcommonword

import (
	"strings"
	"unicode"
)

func mostCommonWord(s string, b []string) string {
	maxCnt := 0
	maxCntWord := ""
	cnts := map[string]int{}

	banned := map[string]struct{}{}

	for i := range b {
		banned[b[i]] = struct{}{}
	}

	n := len(s)

	for l, r := 0, 0; l < n; {

		if !unicode.IsLetter(rune(s[l])) {
			l++
			continue
		}

		r = l

		for r < n && unicode.IsLetter(rune(s[r])) {
			r++
		}

		w := strings.ToLower(s[l:r])
		if _, has := banned[w]; !has {
			cnts[w]++
			if cnts[w] > maxCnt {
				maxCnt = cnts[w]
				maxCntWord = w
			}
		}
		l = r
	}

	return maxCntWord
}

func mostCommonWord1(s string, b []string) string {
	maxCnt := 0
	maxCntWord := ""
	cnts := map[string]int{}

	banned := map[string]struct{}{}

	for i := range b {
		banned[b[i]] = struct{}{}
	}

	n := len(s)

	var word []byte
	for i, n := 0, n; i <= n; i++ {
		if i < n && unicode.IsLetter(rune(s[i])) {
			word = append(word, byte(unicode.ToLower(rune(s[i]))))
		} else if word != nil {
			w := string(word)
			if _, has := banned[w]; !has {
				cnts[w]++
				if cnts[w] > maxCnt {
					maxCnt = cnts[w]
					maxCntWord = w
				}
			}
			word = nil
		}
	}

	return maxCntWord
}
