/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func leftBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if r < len(nums) && nums[r] == target {
		return r
	}
	return -1
}

func rightBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if target >= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l > 0 && nums[l-1] == target {
		return l - 1
	}
	return -1
}

// @lc code=end
