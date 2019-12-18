/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		cnt++
		if cnt > 1 {
			return false
		}
		if (i > 1 && nums[i] < nums[i-2]) && (i < len(nums)-1 && nums[i+1] < nums[i-1]) {
			return false
		}
	}
	return true
}

// @lc code=end
