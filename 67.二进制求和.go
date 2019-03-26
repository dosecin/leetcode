/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (47.14%)
 * Total Accepted:    19.5K
 * Total Submissions: 41.4K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */
func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	var ret []byte
	var addStr string
	if lenA > lenB {
		ret = make([]byte, lenA+1, lenA+1)
		copy(ret[1:], a)
		addStr = b
	} else {
		ret = make([]byte, lenB+1, lenB+1)
		copy(ret[1:], b)
		addStr = a
	}

	ret[0] = '0'
	lenRet := len(ret)
	for i := len(addStr) - 1; i >= 0; i-- {
		lenRet--
		ret[lenRet] += addStr[i] - '0'
		if ret[lenRet] > '1' {
			ret[lenRet] -= 2
			ret[lenRet-1]++
		}
	}

	for lenRet > 0 {
		if ret[lenRet] > '1' {
			ret[lenRet] -= 2
			ret[lenRet-1]++
		}
		lenRet--
	}

	if ret[0] > '0' {
		return string(ret)
	}
	return string(ret[1:])
}
