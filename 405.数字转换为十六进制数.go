/*
 * @lc app=leetcode.cn id=405 lang=golang
 *
 * [405] 数字转换为十六进制数
 */

var hexMap = map[uint32]byte{
	0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7',
	8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f',
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	un := uint32(num)
	b := []byte{}
	for un != 0 {
		n := un & 0xf
		b = append(b, hexMap[n])
		un >>= 4
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

