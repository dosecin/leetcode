/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */
func constructRectangle(area int) []int {
	w, l := 1, area
	res := []int{l, w}
	for w <= l {
		a := w * l
		if a == area {
			res[0] = l
			res[1] = w
			w++
			l--
		} else if a > area {
			l--
		} else {
			w++
		}
	}
	return res
}

