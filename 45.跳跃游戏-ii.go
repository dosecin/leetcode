/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	step, currMax, max := 0, 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if i > currMax {
			step++
			currMax = max
			if currMax == len(nums)-1 {
				break
			}
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return step
}

// @lc code=end
