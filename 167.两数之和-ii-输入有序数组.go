/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	lef, rig := 0, len(numbers)-1
	for lef < rig {
		sum := numbers[lef] + numbers[rig]
		if sum == target {
			return []int{lef + 1, rig + 1}
		}
		if sum > target {
			rig--
		} else {
			lef++
		}
	}
	return []int{}
}

// @lc code=end
