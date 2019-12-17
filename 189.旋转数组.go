/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	reverse := func(nums []int) {
		lef, rig := 0, len(nums)-1
		for lef < rig {
			nums[lef], nums[rig] = nums[rig], nums[lef]
			lef++
			rig--
		}
	}
	k %= len(nums)
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}

// @lc code=end
