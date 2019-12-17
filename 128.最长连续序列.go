/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	numMap := map[int]int{}
	maxLen := 0
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			continue
		}
		left, _ := numMap[num-1]
		right, _ := numMap[num+1]
		length := 1 + left + right
		numMap[num] = length
		numMap[num-left] = length
		numMap[num+right] = length
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

// @lc code=end
