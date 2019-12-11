/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	t, b := 0, len(matrix)-1
	for t < b {
		for i := t; i < b; i++ {
			matrix[t][i], matrix[i][b], matrix[b][b-i+t], matrix[b-i+t][t] =
				matrix[b-i+t][t], matrix[t][i], matrix[i][b], matrix[b][b-i+t]
		}
		t++
		b--
	}
}

// @lc code=end
