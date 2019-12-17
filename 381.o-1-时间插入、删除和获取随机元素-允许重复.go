/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 */

// @lc code=start
import (
	"math/rand"
	"time"
)

type RandomizedCollection struct {
	mapVal map[int]map[int]struct{}
	values []int
	r      *rand.Rand
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		mapVal: map[int]map[int]struct{}{},
		values: []int{},
		r:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	res := true
	if _, ok := this.mapVal[val]; ok {
		res = false
	} else {
		this.mapVal[val] = map[int]struct{}{}
	}
	this.mapVal[val][len(this.values)] = struct{}{}
	this.values = append(this.values, val)
	return res
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	var valIdxs map[int]struct{}
	ok := false
	if valIdxs, ok = this.mapVal[val]; !ok {
		return false
	}
	valIdx := -1
	for valIdx = range valIdxs {
		break
	}
	delete(valIdxs, valIdx)
	if len(valIdxs) == 0 {
		delete(this.mapVal, val)
	}
	lastIdx := len(this.values) - 1
	lastVal := this.values[lastIdx]
	this.values[valIdx], this.values[lastIdx] = lastVal, val
	if lastIdx != valIdx {
		delete(this.mapVal[lastVal], lastIdx)
		this.mapVal[lastVal][valIdx] = struct{}{}
	}
	this.values = this.values[:lastIdx]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	if len(this.values) == 0 {
		return 0
	}
	rd := this.r.Int()
	return this.values[rd%len(this.values)]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
