/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第N个数字
 */
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}

	k, p := 9, 1
	for n > k*p {
		n -= k * p
		k *= 10
		p++
	}
	j := p - 1 - (n-1)%p
	n = k/9 + (n-1)/p
	for j > 0 {
		n /= 10
		j--
	}
	return n % 10
}

