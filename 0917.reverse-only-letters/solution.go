package reverseonlyletters

func reverseOnlyLetters(s string) string {
	arr := []byte(s)

	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		// 非大小写跳过
		for l < r && !((arr[l] >= 'a' && arr[l] <= 'z') || (arr[l] >= 'A' && arr[l] <= 'Z')) {
			l++
		}
		// 非大小写跳过
		for l < r && !((arr[r] >= 'a' && arr[r] <= 'z') || (arr[r] >= 'A' && arr[r] <= 'Z')) {
			r--
		}

		arr[l], arr[r] = arr[r], arr[l]
	}

	return string(arr)
}
