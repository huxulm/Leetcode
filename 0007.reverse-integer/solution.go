package reverseinteger

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < (-1<<31)/10 || rev > (1>>31-1)/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
