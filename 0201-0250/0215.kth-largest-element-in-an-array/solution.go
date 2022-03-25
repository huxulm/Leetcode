package kthlargestelementinanarray

import (
	"container/heap"
	"sort"
)

// 方法一: 快速排序
// 时间: O(nlog(n))
func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]
}

// 大根堆
type hp []int

func (x hp) Len() int           { return len(x) }
func (x hp) Less(i, j int) bool { return x[i] > x[j] }
func (x hp) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x *hp) Push(v interface{}) {
	*x = append(*x, v.(int))
}

func (x *hp) Pop() interface{} {
	old := *x
	n := len(old)
	v := old[n-1]
	*x = old[0 : n-1]
	return v
}

func findKthLargest1(nums []int, k int) int {
	h := hp(nums)
	heap.Init(&h)

	for i := 0; i < k-1; i++ {
		heap.Pop(&h)
	}

	return heap.Pop(&h).(int)
}

// 方法三: 堆（手动）
func findKthLargest2(nums []int, k int) int {
	heapify(nums)
	// fmt.Println(nums)
	for i := 0; i < k-1; i++ {
		pop(&nums)
	}
	return nums[0]
}

// 构建堆
func heapify(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		sink(nums, i, n)
	}
}

// 堆下沉
func sink(nums []int, i, n int) {
	parent := i
	for {
		// 此时使用了nums[0] 所以左子节点为2n+1,右子节点为2n+2
		left := 2*parent + 1
		if left >= n || left < 1 {
			break
		}
		max := left
		//  判断左右子节点的大小
		if right := left + 1; right < n && nums[right] > nums[left] {
			max = right
		}
		if nums[max] < nums[parent] {
			break
		}
		nums[max], nums[parent] = nums[parent], nums[max]
		parent = max
	}
}

func pop(nums *[]int) int {
	rst := (*nums)[0]
	n := len(*nums)
	(*nums)[0], (*nums)[n-1] = (*nums)[n-1], (*nums)[0]
	sink(*nums, 0, n-1)
	(*nums) = (*nums)[:n-1]
	return rst
}

// 方法四: 堆排序（LC官方）
func findKthLargest3(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
