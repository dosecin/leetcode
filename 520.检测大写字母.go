/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */
func detectCapitalUse(word string) bool {
	flag := 0
	for i := range word {
		c := word[i]
		if 'a' <= c && c <= 'z' {
			flag |= 1
		} else if 'A' <= c && c <= 'Z' {
			if i == 0 {
				flag |= 2
			} else {
				flag |= 4
			}
		}
	}
	return flag == 1 || flag == 4 || flag == 3 || flag == 6 || flag == 2
}

