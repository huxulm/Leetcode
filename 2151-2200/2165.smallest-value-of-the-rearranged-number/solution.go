package smallestvalueoftherearrangednumber

import (
	"fmt"
	"sort"
)

func smallestNumber(num int64) int64 {
	nums := []int{}
	isPositive := num > 0

	if !isPositive {
		num = -num
	}

	for num != 0 {
		nums = append(nums, int(num%10))
		num /= 10
	}

	if isPositive {
		sort.Ints(nums)
		if nums[0] == 0 {
			for i := range nums {
				if nums[i] != 0 {
					nums[0], nums[i] = nums[i], nums[0]
					break
				}
			}
		}
	} else {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] > nums[j]
		})
	}

	fmt.Println(nums)

	for i := range nums {
		num = num*10 + int64(nums[i])
	}

	if !isPositive {
		num = -num
	}

	return num
}
