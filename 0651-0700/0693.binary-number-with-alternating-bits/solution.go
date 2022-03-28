package binarynumberwithalternatingbits

func hasAlternatingBits(n int) bool {
	for i := 1; n>>i > 0; i++ {
		if (n>>i&1)^((n>>(i-1))&1) != 1 {
			return false
		}
	}
	return true
}

// 1010101010101 => 0101010101010
// 101010101010  =>  010101010101

func hasAlternatingBits1(n int) bool {
	a := n ^ n>>1
	return a&(a+1) == 0
}
