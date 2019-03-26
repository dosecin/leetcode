/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (48.61%)
 * Total Accepted:    24.6K
 * Total Submissions: 50.6K
 * Testcase Example:  '1'
 *
 * 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 被读作  "one 1"  ("一个一") , 即 11。
 * 11 被读作 "two 1s" ("两个一"）, 即 21。
 * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
 *
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
 *
 * 注意：整数顺序将表示为一个字符串。
 *
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "1"
 *
 *
 * 示例 2:
 *
 * 输入: 4
 * 输出: "1211"
 *
 *
 */

var casMap map[int]string

func init() {
	casMap = make(map[int]string, 30)
	casMap[1] = "1"
	for i := 1; i < 30; i++ {
		str := casMap[i]
		cas := make([]byte, 0)
		cnt := byte(1)
		c := str[0]
		for j := 1; j < len(str); j++ {
			if c != str[j] {
				cas = append(cas, '0'+cnt, c)
				cnt = byte(1)
				c = str[j]
				continue
			}
			cnt++
		}
		cas = append(cas, '0'+cnt, c)
		casMap[i+1] = string(cas)
	}
}

func countAndSay(n int) string {
	return casMap[n]
}
