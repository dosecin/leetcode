/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (60.62%)
 * Total Accepted:    19K
 * Total Submissions: 31.2K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 */
func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	ret := make([][]int, numRows, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = ret[i-1][j-1] + ret[i-1][j]
		}
		ret[i] = row
	}
	return ret
}
