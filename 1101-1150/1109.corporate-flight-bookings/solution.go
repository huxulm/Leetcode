package corporateflightbookings

// 方法一: 树状数组 区间修改+单点查询
func corpFlightBookings(bookings [][]int, n int) (ans []int) {
	var a = make([]int, n+1)
	tree := make([]int, n+1)

	add := func(i, v int) {
		for i < len(tree) {
			tree[i] += v
			i += i & -i
		}
	}

	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}

	// i >= 1
	queryOne := func(i int) int { return a[i] + sum(i) }

	addRange := func(l, r int, val int) { add(l, val); add(r+1, -val) } // [l,r]

	for i := range bookings {
		l, r, val := bookings[i][0], bookings[i][1], bookings[i][2]
		addRange(l, r, val)
	}

	ans = make([]int, n)
	for i := range ans {
		ans[i] = queryOne(i + 1)
	}
	return
}

// 差分
func corpFlightBookings1(bookings [][]int, n int) (ans []int) {
	return
}
