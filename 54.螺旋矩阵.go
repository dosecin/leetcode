/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 {
		return res
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for up <= down && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		for i := up + 1; i < down; i++ {
			res = append(res, matrix[i][right])
		}
		if up != down {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
		}
		if right != left {
			for i := down - 1; i > up; i-- {
				res = append(res, matrix[i][left])
			}
		}
		up++
		down--
		left++
		right--
	}
	return res
}

// @lc code=end
