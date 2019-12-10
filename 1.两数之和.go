/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		diff := target - num
		if j, ok := numsMap[diff]; ok {
			return []int{j, i}
		}
		numsMap[num] = i
	}
	return nil
}

// @lc code=end
