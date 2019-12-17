/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	lef, rig := 0, len(nums)-1
	for lef < rig {
		mid := lef + (rig-lef)/2
		if nums[mid] > nums[rig] {
			lef = mid + 1
		} else if nums[mid] < nums[rig] {
			rig = mid
		} else {
			rig--
		}
	}
	return nums[lef]
}

// @lc code=end
