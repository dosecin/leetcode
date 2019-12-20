/*
 * @lc app=leetcode.cn id=748 lang=golang
 *
 * [748] 最短完整词
 */

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	lpArr := [26]byte{}
	lpCnt := 0
	for _, c := range []byte(licensePlate) {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c >= 'a' && c <= 'z' {
			lpArr[c-'a']++
			lpCnt++
		}
	}
	var ans string
	for _, word := range words {
		wdArr := [26]byte{}
		copy(wdArr[:], lpArr[:])
		wdCnt := lpCnt
		for _, c := range []byte(word) {
			if wdArr[c-'a'] > 0 {
				wdArr[c-'a']--
				wdCnt--
				if wdCnt == 0 {
					break
				}
			}
		}
		if wdCnt == 0 && (ans == "" || len(ans) > len(word)) {
			ans = word
		}
	}
	return ans
}

// @lc code=end
