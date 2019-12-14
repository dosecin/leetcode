/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	rowIndex++
	ans := make([]int, rowIndex, rowIndex)
	for i := 0; i < rowIndex; i++ {
		if i == 0 {
			ans[0] = 1
			continue
		}
		pre := ans[0]
		for j := 0; j <= i; j++ {
			left := 0
			if j > 0 {
				left = pre
			}
			right := 0
			if j < i {
				right = ans[j]
			}
			pre = ans[j]
			ans[j] = left + right
		}
	}
	return ans
}

// @lc code=end
