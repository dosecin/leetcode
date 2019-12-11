/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end
