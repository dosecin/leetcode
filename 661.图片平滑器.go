/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] 图片平滑器
 */

// @lc code=start
func imageSmoother(M [][]int) [][]int {
	m, n := len(M), len(M[0])
	first := make([]int, n)
	second := make([]int, n)
	direct := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1},
		{0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	for i := 0; i < m; i++ {
		if i > 1 {
			k := i - 2
			newData := first
			for j := 0; j < n; j++ {
				M[k][j] = newData[j]
			}
		}
		for j := 0; j < n; j++ {
			cnt, sum := 1, M[i][j]
			for _, d := range direct {
				ii, jj := i+d[0], j+d[1]
				if ii < 0 || ii >= m {
					continue
				}
				if jj < 0 || jj >= n {
					continue
				}
				cnt++
				sum += M[ii][jj]
			}
			first[j] = sum / cnt
		}
		first, second = second, first
	}
	if m > 1 {
		k := m - 2
		newData := first
		for j := 0; j < n; j++ {
			M[k][j] = newData[j]
		}
	}
	if m > 0 {
		k := m - 1
		newData := second
		for j := 0; j < n; j++ {
			M[k][j] = newData[j]
		}
	}
	return M
}

// @lc code=end
