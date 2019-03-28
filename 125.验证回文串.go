/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (38.62%)
 * Total Accepted:    28.8K
 * Total Submissions: 74.6K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */
func isPalindrome(s string) bool {
	head, tail := 0, len(s)-1
	for head < tail {
		ch := convertC(s[head])
		if skip(ch) {
			head++
			continue
		}
		ct := convertC(s[tail])
		if skip(ct) {
			tail--
			continue
		}
		if ch != ct {
			return false
		}
		head++
		tail--
	}
	return true
}

func convertC(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return 'A' + c - 'a'
	}
	return c
}

func skip(c byte) bool {
	if '0' <= c && c <= '9' {
		return false
	}
	if 'A' <= c && c <= 'Z' {
		return false
	}
	return true
}
