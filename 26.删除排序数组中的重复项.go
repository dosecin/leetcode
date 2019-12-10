/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	pre, curr := 0, 1
	for curr < len(nums) {
		if nums[curr] == nums[curr-1] {
			curr++
			continue
		}
		pre++
		nums[pre] = nums[curr]
		curr++
	}
	return pre + 1
}

// @lc code=end
