package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			return []int{l, r}
		}
	}
	return []int{-1, -1}
}
