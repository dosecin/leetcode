/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */
func getSum(a int, b int) int {	
	for b != 0 {
		a, b = a ^ b, (a&b)<<1
	}
	return a
}

