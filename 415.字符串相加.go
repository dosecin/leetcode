/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */
func addStrings(num1 string, num2 string) string {
    if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	carry := false
	retlen := 5100
	ret := make([]byte, retlen, retlen)
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 {
		retlen--
		ret[retlen] = num1[i] + num2[j] - '0'
		if carry {
			ret[retlen]++
			carry = false
		}
		if ret[retlen] > '9' {
			ret[retlen] -= 10
			carry = true
		}
		i--
		j--
	}
	for j >= 0 {
		retlen--
		ret[retlen] = num2[j]
		if carry {
			ret[retlen]++
			carry = false
		}
		if ret[retlen] > '9' {
			ret[retlen] -= 10
			carry = true
		}
		j--
	}
	if carry {
		retlen--
		ret[retlen] = '1'
	}
	return string(ret[retlen:])
}

