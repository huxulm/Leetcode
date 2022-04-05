package primenumberofsetbitsinbinaryrepresentation

import "math/bits"

// 埃氏筛
func countPrimeSetBits(left int, right int) (ans int) {
	//  isPrime[i] 表示数 i 是不是质数
	n := 32
	isPrime := make([]bool, n)
	for i := range isPrime[2:] {
		isPrime[i+2] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := left; i <= right; i++ {
		if isPrime[bits.OnesCount(uint(i))] {
			ans++
		}
	}
	return
}
