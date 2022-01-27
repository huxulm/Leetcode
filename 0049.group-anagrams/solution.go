package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for i := range strs {
		code := encodeBySort(strs[i])
		if _, ok := res[code]; !ok {
			res[code] = []string{}
		}
		res[code] = append(res[code], strs[i])
	}
	ans := make([][]string, 0, len(res))
	for _, v := range res {
		ans = append(ans, v)
	}
	return ans
}

func encode(s string) string {
	code := make([]byte, 26)
	for i := range s {
		delta := s[i] - 'a'
		code[delta]++
	}
	return string(code)
}

func encodeBySort(s string) string {
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}
