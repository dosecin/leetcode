/*
 * @lc app=leetcode.cn id=521 lang=golang
 *
 * [521] 最长特殊序列 Ⅰ
 */
func findLUSlength(a string, b string) int {
    if strings.Compare(a, b) == 0 {
		return -1
	}
	al := len(a)
	bl := len(b)
	if al > bl {
		return al
	}
	return bl
}

