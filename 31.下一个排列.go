/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			reverse(nums, i, len(nums)-1)
			idx := i - 1
			for i < len(nums) {
				if nums[i] > nums[idx] {
					break
				}
				i++
			}
			nums[idx], nums[i] = nums[i], nums[idx]
			return
		}
	}
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// @lc code=end
