package reversebits

func reverseBits(n uint32) (rev uint32) {

	for i := 0; i < 32; i++ {
		rev |= (n & 1) << (32 - i)
		n >>= 1
	}

	return
}
