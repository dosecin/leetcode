/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	cnt, preIdx := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			nums[preIdx] = nums[i]
			preIdx++
		}
	}
	return preIdx
}

// @lc code=end
