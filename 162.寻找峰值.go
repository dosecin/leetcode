/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	maxIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			maxIdx = i
		} else {
			break
		}
	}
	return maxIdx
}

// @lc code=end
