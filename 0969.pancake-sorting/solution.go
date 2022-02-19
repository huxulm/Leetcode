package pancakesorting

// 冒泡排序: 每一轮将最大值反转到最后
func pancakeSort(arr []int) (ans []int) {
	var reverse = func(arr []int) {
		n := len(arr)
		for i := 0; i < n>>1; i++ {
			arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
		}
	}

	n := len(arr)

	for i := n - 1; i > 0; i-- {
		idx := 0
		for j := 0; j <= i; j++ {
			if arr[j] > arr[idx] {
				idx = j
			}
		}
		// 煎饼排序, 将arr[idx]翻转到[i]位置
		if idx == i {
			continue
		}
		reverse(arr[:idx+1])
		reverse(arr[:i+1])
		ans = append(ans, idx+1, i+1)
	}

	return
}

func pancakeSort1(arr []int) (ans []int) {
	for n := len(arr); n > 1; n-- {
		index := 0
		for i, v := range arr[:n] {
			if v > arr[index] {
				index = i
			}
		}
		if index == n-1 {
			continue
		}
		for i, m := 0, index; i < (m+1)/2; i++ {
			arr[i], arr[m-i] = arr[m-i], arr[i]
		}
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		ans = append(ans, index+1, n)
	}
	return
}
