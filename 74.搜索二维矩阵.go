/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := l + (r-l)>>1
		i, j := mid/n, mid%n
		if target > matrix[i][j] {
			l = mid + 1
		} else if target < matrix[i][j] {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}

// @lc code=end
