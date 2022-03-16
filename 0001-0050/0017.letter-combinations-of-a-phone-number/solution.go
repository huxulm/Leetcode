package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	m := map[byte][]byte{
		// '0': []byte{},
		// '1': []byte{},
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	res := make([][]byte, 0)

	backtrack(m, &res, &digits, 0, make([]byte, len(digits)))

	arr := make([]string, len(res))
	for i := range res {
		arr[i] = string(res[i])
	}
	return arr
}

// 选择第i个数字对应字母
func backtrack(table map[byte][]byte, result *[][]byte, s *string, index int, track []byte) {
	if index == len(*s) {
		return
	}
	list := table[(*s)[index]]
	isLast := len(*s)-1 == index
	for i := range list {
		// 选择
		track[index] = list[i]
		if isLast {
			res := make([]byte, len(track))
			copy(res, track)
			*result = append(*result, res)
		} else {
			// 选择下一个
			backtrack(table, result, s, index+1, track)
		}
	}
}
