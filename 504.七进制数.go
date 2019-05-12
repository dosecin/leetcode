/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	b := []byte{}
	neg := false
	if num < 0 {
		neg = true
		num = -num
	}
	for num != 0 {
		b = append(b, byte(num%7)+'0')
		num /= 7
	}
	if neg {
		b = append(b, '-')
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

