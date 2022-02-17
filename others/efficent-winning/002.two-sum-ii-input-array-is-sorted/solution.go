package twosumiiinputarrayissorted

import "sort"

func twoSum(numbers []int, target int) []int {

	sort.Ints(numbers)

	l, r := 0, len(numbers)-1

	for l < r {
		temp := numbers[l] + numbers[r]
		if temp == target {
			return []int{l + 1, r + 1}
		} else if temp < target {
			l++
		} else {
			r--
		}
	}

	return []int{-1, -1}
}

// 双指针+二分查找
func twoSum1(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {
		lo, hi := i+1, len(numbers)-1
		for lo <= hi {
			mid := (lo + hi) >> 1
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}

	return []int{-1, -1}
}
