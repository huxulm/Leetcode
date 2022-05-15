package p004

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		tot := int32(0)
		for _, x := range nums {
			tot += int32(x) >> i & 1
		}
		// 多的一位是 ans 提供的
		if tot%3 != 0 {
			ans |= (1 << i)
		}
	}
	return int(ans)
}
