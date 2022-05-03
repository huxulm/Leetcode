package reorderdatainlogfiles

import "sort"

func reorderLogFiles(logs []string) (ans []string) {
	alphas := [][]string{}
	nums := []string{}

	for _, log := range logs {
		var isAlpha, isNum bool = true, true
		var start int = -1 // content start index
		for i, ch := range log {
			if start == -1 {
				if ch == ' ' {
					start = i
				}
			} else {
				if ch == ' ' {
					continue
				}
				if ch >= '0' && ch <= '9' {
					isAlpha = false
					break
				} else if ch >= 'a' && ch <= 'z' {
					isNum = false
					break
				}
			}
		}
		if isAlpha {
			alphas = append(alphas, []string{log[:start], log[start:]})
		} else if isNum {
			nums = append(nums, log)
		}
	}
	sort.Slice(alphas, func(i, j int) bool {
		if alphas[i][1] != alphas[j][1] {
			return alphas[i][1] < alphas[j][1]
		} else {
			return alphas[i][0] < alphas[j][0]
		}
	})

	for i := range alphas {
		ans = append(ans, alphas[i][0]+alphas[i][1])
	}
	for i := range nums {
		ans = append(ans, nums[i])
	}

	return
}
