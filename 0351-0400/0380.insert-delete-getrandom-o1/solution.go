package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	indexes map[int]int
	nums    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{indexes: make(map[int]int), nums: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexes[val]; !ok {
		this.indexes[val] = len(this.nums)
		this.nums = append(this.nums, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.indexes[val]; ok {
		last := len(this.nums) - 1
		this.nums[last], this.nums[index] = this.nums[index], this.nums[last]
		this.indexes[this.nums[index]] = index
		this.nums = this.nums[:last]
		delete(this.indexes, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
