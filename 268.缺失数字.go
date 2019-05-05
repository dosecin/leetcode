/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */
func missingNumber(nums []int) int {
	var num int
	for i := range nums {
		num += i - nums[i]
	}
	return num + len(nums)
}
