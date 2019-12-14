/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	type flag struct {
		left, right, height int
	}

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	calcArea := func(f flag) int {
		return f.height * (f.right - f.left + 1)
	}
	m, n := len(matrix), len(matrix[0])
	maxArea := 0
	flags := []flag{}
	for i := 0; i < m; i++ {
		// 更新高度
		for j := 0; j < n; j++ {
			if j >= len(flags) {
				f := flag{right: n}
				if matrix[i][j] == '1' {
					f.height = 1
				}
				flags = append(flags, f)
			} else {
				if matrix[i][j] == '1' {
					flags[j].height++
				} else {
					flags[j].height = 0
				}
			}
		}
		// 更新left
		currLeft := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if currLeft > flags[j].left {
					flags[j].left = currLeft
				}
			} else {
				flags[j].left = 0
				currLeft = j + 1
			}
		}
		// 更新right
		currRight := n - 1
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				if currRight < flags[j].right {
					flags[j].right = currRight
				}
			} else {
				flags[j].right = n
				currRight = j - 1
			}
		}
		// 更新area
		for j := 0; j < n; j++ {
			area := calcArea(flags[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// @lc code=end
