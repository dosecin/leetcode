/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */
func findTheDifference(s string, t string) byte {
	n := [26]int{}
	for _, b := range s {
		n[b-'a']++
	}
	for _, b := range t {
		n[b-'a']--
	}
	for i := range n {
		if n[i] != 0 {
			return byte(i)+'a'
		}
	}
	return 0
}

