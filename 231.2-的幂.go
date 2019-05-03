/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
    for n & 1 == 0 {
		n >>= 1
	}
	return n == 1
}

