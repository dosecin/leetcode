/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */
func addDigits(num int) int {
	for num > 9 {
		var sum int
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}
	return num
}
