/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (54.11%)
 * Total Accepted:    11.7K
 * Total Submissions: 21.6K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */
func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		ret[0], ret[i] = 1, 1
		l := ret[0]
		for j := 1; j < i; j++ {
			r := ret[j]
			ret[j] = l + r
			l = r
		}
	}
	return ret
}
