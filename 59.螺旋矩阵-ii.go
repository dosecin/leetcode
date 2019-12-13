/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	t, b, l, r := 0, n-1, 0, n-1
	num, max := 1, n*n
	ans := make([][]int, n, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n, n)
	}
	for num <= max {
		for i := l; i <= r; i++ {
			ans[t][i] = num
			num++
		}
		for i := t + 1; i <= b; i++ {
			ans[i][r] = num
			num++
		}
		if t != b {
			for i := r - 1; i > l; i-- {
				ans[b][i] = num
				num++
			}
		}
		if l != r {
			for i := b; i > t; i-- {
				ans[i][l] = num
				num++
			}
		}

		t++
		b--
		l++
		r--
	}
	return ans
}

// @lc code=end
