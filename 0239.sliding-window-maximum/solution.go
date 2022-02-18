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

type MontonicQueue struct {
	data []int
}

func (mq *MontonicQueue) Push(n int) {
	// 删除尾部所有小于n的元素
	for sz := len(mq.data); sz > 0 && n > mq.data[sz-1]; sz = len(mq.data) {
		mq.data = mq.data[:sz-1]
	}
	mq.data = append(mq.data, n)
}

func (mq *MontonicQueue) Max() int {
	return mq.data[0]
}

func (mq *MontonicQueue) Pop(n int) {
	if mq.data[0] == n {
		mq.data = mq.data[1:]
	}
}

func maxSlidingWindow(nums []int, k int) (ans []int) {

	n := len(nums)

	if k == 1 {
		return nums
	}

	// 使用单调队列实现
	// window := &MontonicQueue{deque: list.New()}
	window := new(MontonicQueue)

	for i := 0; i < n; i++ {
		if i < k-1 {
			window.Push(nums[i])
		} else {
			window.Push(nums[i])
			ans = append(ans, window.Max())
			window.Pop(nums[i-k+1])
		}
	}

	return ans
}
