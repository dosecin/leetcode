/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	lastIdx := len(nums) - 1
	lef, rig := 0, -1
	max, min := nums[lef], nums[lastIdx]
	for i := range nums {
		if nums[i] < max {
			rig = i
		} else {
			max = nums[i]
		}
		if nums[lastIdx-i] > min {
			lef = lastIdx - i
		} else {
			min = nums[lastIdx-i]
		}
	}
	return rig - lef + 1
}

// @lc code=end
