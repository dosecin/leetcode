/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	lef, rig := 0, len(nums)-1
	for lef < rig {
		mid := lef + (rig-lef)/2
		if nums[mid] > nums[rig] {
			lef = mid + 1
		} else {
			rig = mid
		}
	}
	return nums[lef]
}

// @lc code=end
