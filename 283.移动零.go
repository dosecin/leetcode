/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	pre, cur := 0, 0
	for cur < len(nums) {
		if pre == cur {
			if nums[cur] != 0 {
				pre++
			}
			cur++
		} else {
			if nums[cur] != 0 {
				nums[pre], nums[cur] = nums[cur], 0
				pre++
			}
			cur++
		}
	}
}

// @lc code=end
