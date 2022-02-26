package numberof1bits

func hammingWeight(num uint32) (ans int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ans++
		}
	}
	return
}

func hammingWeight1(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}
