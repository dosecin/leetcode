/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	lef, rig := 0, len(nums)-1
	for lef <= rig {
		mid := lef + (rig-lef)>>1
		if target == nums[mid] {
			return true
		}
		if nums[lef] == nums[mid] { // 无法区分有序部分
			lef++
			continue
		}
		if nums[lef] < nums[mid] { // 前半部分有序
			if target >= nums[lef] && target < nums[mid] { //位于有序部分
				rig = mid - 1
			} else {
				lef = mid + 1
			}
		} else { // 后半部分有序
			if nums[mid] < target && target <= nums[rig] { //位于有序部分
				lef = mid + 1
			} else {
				rig = mid - 1
			}
		}
	}
	return false
}

// @lc code=end
