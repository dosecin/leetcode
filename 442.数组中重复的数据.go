/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	ans := []int{}
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		num--
		if nums[num] < 0 {
			ans = append(ans, num+1)
		}
		nums[num] = -nums[num]
	}
	return ans
}

// @lc code=end
