package slidingwindowmaximum

// type MontonicQueue struct {
// 	deque *list.List
// }

// // 在队尾添加元素 n
// func (mq *MontonicQueue) Push(n int) {
// 	for mq.deque.Len() > 0 && mq.deque.Back().Value.(int) < n {
// 		mq.deque.Remove(mq.deque.Back())
// 	}
// 	mq.deque.PushBack(n)
// }

// // 返回当前队列中的最大值
// func (mq *MontonicQueue) Max() int {
// 	return mq.deque.Front().Value.(int)
// }

// // 队头元素如果是 n，删除它
// func (mq *MontonicQueue) Pop(n int) {
// 	if mq.deque.Len() > 0 && mq.deque.Front().Value.(int) == n {
// 		mq.deque.Remove(mq.deque.Front())
// 	}
// }

// type MontonicQueue struct {
// 	data []int
// }

// func (mq *MontonicQueue) Push(n int) {
// 	// 删除尾部所有小于n的元素
// 	for sz := len(mq.data); sz > 0 && n > mq.data[sz-1]; sz = len(mq.data) {
// 		mq.data = mq.data[:sz-1]
// 	}
// 	mq.data = append(mq.data, n)
// }

// func (mq *MontonicQueue) Max() int {
// 	return mq.data[0]
// }

// func (mq *MontonicQueue) Pop(n int) {
// 	if mq.data[0] == n {
// 		mq.data = mq.data[1:]
// 	}
// }

// func maxSlidingWindow(nums []int, k int) (ans []int) {

// 	n := len(nums)

// 	if k == 1 {
// 		return nums
// 	}

// 	// 使用单调队列实现
// 	// window := &MontonicQueue{deque: list.New()}
// 	window := new(MontonicQueue)

// 	for i := 0; i < n; i++ {
// 		if i < k-1 {
// 			window.Push(nums[i])
// 		} else {
// 			window.Push(nums[i])
// 			ans = append(ans, window.Max())
// 			window.Pop(nums[i-k+1])
// 		}
// 	}

// 	return ans
// }

type MontonicQueue struct {
	data []int
}

func (mq *MontonicQueue) Len() int {
	return len(mq.data)
}

func (mq *MontonicQueue) Push(x int) {
	// 从尾部加入, 移除队列中所有小于x的元素
	for mq.Len() > 0 && mq.Back() < x {
		mq.RemoveBack()
	}
	mq.data = append(mq.data, x)
}

func (mq *MontonicQueue) Pop(x int) {
	if mq.Len() > 0 && mq.Front() == x {
		mq.RemoveFront()
	}
}

func (mq *MontonicQueue) Front() int {
	return mq.data[0]
}
func (mq *MontonicQueue) Back() int {
	return mq.data[mq.Len()-1]
}

func (mq *MontonicQueue) RemoveFront() {
	mq.data = mq.data[1:mq.Len()]
}

func (mq *MontonicQueue) RemoveBack() {
	mq.data = mq.data[0 : mq.Len()-1]
}

func (mq *MontonicQueue) Max() int {
	return mq.Front()
}

// 求窗口中的最大值
func maxSlidingWindow(nums []int, k int) (ans []int) {
	n := len(nums)
	window := new(MontonicQueue)

	for i := 0; i < n; i++ {
		if i < k {
			window.Push(nums[i])
			if i == k-1 {
				ans = append(ans, window.Max())
			}
		} else {
			window.Pop(nums[i-k])
			window.Push(nums[i])
			ans = append(ans, window.Max())
		}
	}
	return
}

// 单调递增栈
type MontonicIncreaseQueue struct {
	data []int
}

func (mq *MontonicIncreaseQueue) Min() int {
	return mq.data[0]
}

func (mq *MontonicIncreaseQueue) Len() int {
	return len(mq.data)
}

func (mq *MontonicIncreaseQueue) Back() int {
	return mq.data[mq.Len()-1]
}

func (mq *MontonicIncreaseQueue) Front() int {
	return mq.data[0]
}

func (mq *MontonicIncreaseQueue) RemoveFront() {
	mq.data = mq.data[1:mq.Len()]
}

func (mq *MontonicIncreaseQueue) RemoveBack() {
	mq.data = mq.data[:mq.Len()-1]
}

func (mq *MontonicIncreaseQueue) Push(x int) {
	// 移除结尾所有大于x的元素
	for mq.Len() > 0 && mq.Back() > x {
		mq.RemoveBack()
	}
	mq.data = append(mq.data, x)
}

func (mq *MontonicIncreaseQueue) Pop(x int) {
	if mq.Len() > 0 && x == mq.Front() {
		mq.RemoveFront()
	}
}

func minSlidingWindow(nums []int, k int) (ans []int) {
	n := len(nums)
	window := new(MontonicIncreaseQueue)

	for i := 0; i < n; i++ {
		if i < k {
			window.Push(nums[i])
			if i == k-1 {
				ans = append(ans, window.Min())
			}
		} else {
			window.Push(nums[i])
			window.Pop(nums[i-k])
			ans = append(ans, window.Min())
		}
	}
	return
}
