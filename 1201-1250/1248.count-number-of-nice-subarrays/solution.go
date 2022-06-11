package countnumberofnicesubarrays

// 方法一: 数学统计
func numberOfSubarrays(nums []int, k int) (ans int) {
	n := len(nums)
	// 记录奇数索引位置
	odd := make([]int, n+2)
	cnt := 0
	for i := range nums {
		if nums[i]&1 == 1 {
			cnt++
			odd[cnt] = i
		}
	}

	cnt++
	// 边界
	odd[0], odd[cnt] = -1, n
	for i := 1; i+k <= cnt; i++ {
		ans += (odd[i] - odd[i-1]) * (odd[i+k] - odd[i+k-1])
	}
	return
}

func numberOfSubarrays1(nums []int, k int) (ans int) {
	n := len(nums)
	cnt := make([]int, n+1)
	odd := 0
	cnt[0] = 1
	for i := range nums {
		odd += nums[i] & 1
		if odd >= k {
			ans += cnt[odd-k]
		}
		cnt[odd] += 1
	}
	return
}
