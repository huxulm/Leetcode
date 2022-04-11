package intervalsbetweenidenticalelements

func getDistances(arr []int) (ans []int64) {
	n := len(arr)

	// 计算 arr[i] 之前和 arr[i] 相等元素的距离只和
	prefix := make([]int, n)
	// 计算 arr[i] 之后和 arr[i] 相等元素的距离只和
	suffix := make([]int, n)

	var m = map[int]int{}
	var cnt = map[int]int{}

	for i, v := range arr {
		cnt[v]++
		if i == 0 {
			m[v] = i
			continue
		}

		if lasti, ok := m[v]; ok {
			prefix[i] = prefix[lasti] + (i-lasti)*(cnt[v]-1)
		}
		// 更新 lasti
		m[v] = i
	}

	m = map[int]int{}
	cnt = map[int]int{}
	for i := n - 1; i >= 0; i-- {
		v := arr[i]
		cnt[v]++

		if lasti, ok := m[v]; ok {
			suffix[i] = suffix[lasti] + (lasti-i)*(cnt[v]-1)
		}
		// 更新 lasti
		m[v] = i
	}

	ans = make([]int64, n)
	for i := 0; i < n; i++ {
		ans[i] = int64(prefix[i]) + int64(suffix[i])
	}
	return
}
