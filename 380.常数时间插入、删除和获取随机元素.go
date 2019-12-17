/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 */

// @lc code=start
import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	mapVal map[int]int
	values []int
	r      *rand.Rand
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		mapVal: map[int]int{},
		values: []int{},
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mapVal[val]; ok {
		return false
	}
	idx := len(this.values)
	this.values = append(this.values, val)
	this.mapVal[val] = idx
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	valIdx, ok := -1, false
	if valIdx, ok = this.mapVal[val]; !ok {
		return false
	}
	lastIdx := len(this.values) - 1
	lastVal := this.values[lastIdx]
	this.values[valIdx], this.values[lastIdx] = lastVal, val
	delete(this.mapVal, val)
	if lastVal != val {
		this.mapVal[lastVal] = valIdx
	}
	this.values = this.values[:lastIdx]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.values) == 0 {
		return 0
	}
	return this.values[this.r.Intn(len(this.values))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
