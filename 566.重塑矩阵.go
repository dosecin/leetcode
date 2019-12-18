/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) == 0 {
		return nums
	}
	m, n := len(nums), len(nums[0])
	if r*c != m*n {
		return nums
	}
	ans := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j
			ar, ac := idx/c, idx%c
			if ar >= len(ans) {
				ans = append(ans, make([]int, c))
			}
			ans[ar][ac] = nums[i][j]
		}
	}
	return ans
}

// @lc code=end
