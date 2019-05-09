/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 */
func minMoves(nums []int) int {
    if len(nums) < 1 {
		return 0
	}
	sum, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if min > nums[i] {
			min = nums[i]
		}
	}
	return sum - min*len(nums)
}

