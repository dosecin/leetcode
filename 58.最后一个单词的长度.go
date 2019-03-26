/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (29.08%)
 * Total Accepted:    21.1K
 * Total Submissions: 72.5K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */
func lengthOfLastWord(s string) int {
	tail := len(s) - 1
	for tail >= 0 {
		if s[tail] != ' ' {
			break
		}
		tail--
	}

	cnt := 0
	for tail >= 0 {
		if s[tail] == ' ' {
			break
		}
		cnt++
		tail--
	}
	return cnt
}
