package solution

import "container/list"

func permute(nums []int) [][]int {
	var res [][]int
	backtrack(&res, list.New(), nums)
	return res
}

func list2IntArr(l *list.List) []int {
	arr := make([]int, l.Len())
	for i, n := 0, l.Front(); n != nil; i, n = i+1, n.Next() {
		arr[i] = n.Value.(int)
	}
	return arr
}

func backtrack(res *[][]int, track *list.List, nums []int) {
	if track.Len() == len(nums) {
		*res = append(*res, list2IntArr(track))
		return
	}
	// 选择列表
	for i := 0; i < len(nums); i++ {
		// 选择track中没有的
		var used bool
		for n := track.Front(); n != nil; n = n.Next() {
			if n.Value == nums[i] {
				used = true
				break
			}
		}
		if used {
			continue
		}
		// 选择
		track.PushBack(nums[i])
		backtrack(res, track, nums)
		// 取消选择
		track.Remove(track.Back())
	}
}
