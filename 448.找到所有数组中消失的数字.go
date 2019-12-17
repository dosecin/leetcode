/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		num--
		if nums[num] > 0 {
			nums[num] = -nums[num]
		}
	}

	ans := []int{}
	for i, num := range nums {
		if num > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// @lc code=end
