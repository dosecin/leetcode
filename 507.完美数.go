/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	fac := []int{1}
	a, b := 2, num/2
	for a < b {
		if num%a == 0 {
			fac = append(fac, a, b)
		}
		a++
		b = num / a
	}
	sum := 0
	for _, f := range fac {
		sum += f
	}
	return num == sum
}

