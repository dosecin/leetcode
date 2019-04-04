/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (26.76%)
 * Total Accepted:    14.9K
 * Total Submissions: 55.2K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */
func countPrimes(n int) int {
	primes := make([]bool, n, n)
	count := 0
	for i := 2; i < n; {
		count++
		for j := i; j < n; j += i {
			primes[j] = true
		}
		for i < n && primes[i] {
			i++
		}
	}
	return count
}
