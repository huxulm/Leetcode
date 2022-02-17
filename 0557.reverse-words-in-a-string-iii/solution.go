package reversewordsinastringiii

func reverseWords(s string) string {
	var str_arr = []byte(s)

	var reverseWord = func(arr []byte) {
		n := len(arr)
		for i := 0; i < n>>1; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
	}

	// 单词开头，结尾索引
	start, end := 0, -2
	for i := range str_arr {
		if str_arr[i] == ' ' {
			start, end = end+2, i-1
			reverseWord(str_arr[start : end+1])
		}
		if i == len(str_arr)-1 {
			reverseWord(str_arr[end+2:])
		}
	}
	return string(str_arr)
}
