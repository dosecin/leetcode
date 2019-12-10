/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	pre := -1
	for _, v := range nums {
		if v != val {
			pre++
			nums[pre] = v
		}
	}
	return pre + 1
}

// @lc code=end
