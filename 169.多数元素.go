/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	cnt, pre := 0, 0
	for i := range nums {
		if cnt == 0 {
			cnt = 1
			pre = nums[i]
		} else {
			if nums[i] == pre {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return pre
}

// @lc code=end
