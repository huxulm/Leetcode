package randompickindex

import "math/rand"

// 方法一: 二分查找
// type Solution struct {
// 	arr [][]int
// }

// 时间: O(nlog(n)) 排序 O(n) 遍历
// func Constructor(nums []int) Solution {
// 	// sort.Ints(nums)
// 	n := len(nums)
// 	rand.Seed(int64(n))
// 	arr := make([][]int, n)
// 	for i := range nums {
// 		arr[i] = []int{nums[i], i}
// 	}
// 	sort.Slice(arr, func(i, j int) bool { return arr[i][0] < arr[j][0] })
// 	return Solution{arr}
// }

// 时间: O(log(n))查找
// func (this *Solution) Pick(target int) int {
// 	l := sort.Search(len(this.arr), func(i int) bool { return this.arr[i][0] >= target })
// 	r := sort.Search(len(this.arr), func(i int) bool { return this.arr[i][0] >= target+1 }) - 1
// 	if r == l {
// 		return this.arr[l][1]
// 	}
// 	return this.arr[l+rand.Intn(r-l+1)][1]
// }

// 方法二: 哈希分组
// type Solution map[int][]int

// func Constructor(nums []int) Solution {
// 	pos := Solution{}
// 	for i := range nums {
// 		pos[nums[i]] = append(pos[nums[i]], i)
// 	}
// 	return pos
// }

// func (pos Solution) Pick(target int) int {
// 	return pos[target][rand.Intn(len(pos[target]))]
// }

// 方法三: 蓄水池抽样
// https://www.jianshu.com/p/7a9ea6ece2af
// https://www.geeksforgeeks.org/reservoir-sampling/
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (s Solution) Pick(target int) (ans int) {
	n := len(s.nums)
	for i, cnt := 0, 0; i < n; i++ {
		if s.nums[i] == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return ans
}
