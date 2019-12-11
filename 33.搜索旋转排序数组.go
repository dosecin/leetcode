/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[j] {
			if target >= nums[mid] || target < nums[i] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		} else if nums[mid] < nums[i] {
			if target <= nums[mid] || target > nums[j] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else {
			if target < nums[mid] {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}
	return -1
}

// @lc code=end
