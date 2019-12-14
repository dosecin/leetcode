/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	ans := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1, i+1)
		if i == 0 {
			row[0] = 1
		} else {
			up := ans[i-1]
			for j := 0; j <= i; j++ {
				left := 0
				if j > 0 {
					left = up[j-1]
				}
				right := 0
				if j < len(up) {
					right = up[j]
				}
				row[j] = left + right
			}
		}
		ans[i] = row
	}
	return ans
}

// @lc code=end
