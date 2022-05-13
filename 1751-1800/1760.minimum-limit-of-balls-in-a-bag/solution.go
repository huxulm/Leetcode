package minimumlimitofballsinabag

import "sort"

func minimumSize(nums []int, maxOps int) int {
	// 找最小，往左侧缩小搜索范围
	return sort.Search(1e9, func(i int) bool {
		i++
		tot := 0
		for _, v := range nums {
			// 向下取整
			tot += (v - 1) / i
		}
		return tot <= maxOps
	}) + 1
}
