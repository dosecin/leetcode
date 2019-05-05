/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
