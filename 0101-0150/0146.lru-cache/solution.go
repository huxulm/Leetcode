package lrucache

import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start

// 方法: 哈希 + 双向链表
type LRUCache struct {
	// 最近使用的元素放在栈顶
	sta *list.List
	m   map[int]*list.Element // key, idx
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, sta: list.New(), m: map[int]*list.Element{}}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.m[key] = this.sta.PushBack(this.sta.Remove(v))
		return v.Value.([]int)[1]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok { // key 存在, 更新
		delete(this.m, this.sta.Remove(v).([]int)[0])
		this.m[key] = this.sta.PushBack([]int{key, value})
	} else { // key 不存在
		if this.sta.Len() == this.cap {
			kv := this.sta.Remove(this.sta.Front()).([]int)
			delete(this.m, kv[0])
		}
		this.m[key] = this.sta.PushBack([]int{key, value})
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
