/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */
func findComplement(num int) int {
	b, n := 1, uint32(num)
	for n != 0 {
		b <<= 1
		n >>= 1
	}
	b--
	return b&^num
}

