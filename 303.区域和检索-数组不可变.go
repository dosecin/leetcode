/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */
type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1, len(nums)+1)
	for i := range nums {
		preSum[i+1] = preSum[i]+nums[i]
	}
	return NumArray{data:preSum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.data[j+1] - this.data[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
