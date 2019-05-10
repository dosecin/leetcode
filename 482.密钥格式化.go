/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */
func licenseKeyFormatting(S string, K int) string {
	buff := []byte{}
	cnt := K
	for i := len(S) - 1; i >= 0; i-- {
		c := S[i]
		if c == '-' {
			continue
		}
		if 'a' <= c && c <= 'z' {
			c = 'A' + c - 'a'
		}
		buff = append(buff, c)
		cnt--
		if cnt == 0 {
			buff = append(buff, '-')
			cnt = K
		}
	}
	if len(buff) == 0{
		return ""
	}
	for i, j := 0, len(buff)-1; i < j; i, j = i+1, j-1 {
		buff[i], buff[j] = buff[j], buff[i]
	}
	if buff[0] == '-' {
		return string(buff[1:])
	}
	return string(buff)
}

