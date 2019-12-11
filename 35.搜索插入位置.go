/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

// @lc code=end
