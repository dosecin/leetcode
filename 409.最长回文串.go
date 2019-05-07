/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */
func longestPalindrome(s string) int {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}
	sum, hasOne := 0, false
	for _, v := range cnt {
		if v % 2 == 0 {
			sum += v
		} else {
			sum += v-1
			hasOne = true
		}
	}
	if hasOne {
		sum++
	}
	return sum
}

