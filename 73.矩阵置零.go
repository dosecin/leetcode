/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	bRowZero := false
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			bRowZero = true
			break
		}
	}
	bColZero := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			bColZero = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if bRowZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if bColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end
