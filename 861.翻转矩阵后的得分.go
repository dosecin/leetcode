/*
 * @lc app=leetcode.cn id=861 lang=golang
 *
 * [861] 翻转矩阵后的得分
 */

// @lc code=start
func matrixScore(A [][]int) int {
	row, col := len(A), len(A[0])
	for i := 0; i < row; i++ {
		if A[i][0] == 0 {
			for j := 0; j < col; j++ {
				A[i][j] = 1 - A[i][j]
			}
		}
	}
	lmt := row / 2
	for j := 1; j < col; j++ {
		sum := 0
		for i := 0; i < row; i++ {
			if A[i][j] == 1 {
				sum++
			}
		}
		if sum > lmt {
			continue
		}
		for i := 0; i < row; i++ {
			A[i][j] = 1 - A[i][j]
		}
	}
	ans := 0
	for i := 0; i < row; i++ {
		curr := 0
		for j := 0; j < col; j++ {
			curr = (curr << 1) | A[i][j]
		}
		ans += curr
	}
	return ans
}

// @lc code=end
