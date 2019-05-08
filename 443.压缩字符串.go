/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */
func compress(chars []byte) int {
	if len(chars) < 1 {
		return 0
	}
	pre, pos, cnt := chars[0], 1, 0
	for i := 0; i < len(chars); i++ {
		if pre == chars[i] {
			cnt++
			continue
		}
		pre = chars[i]
		if cnt > 1 {
			pos = compressItoa(chars, cnt, pos)
		}
		chars[pos] = pre
		pos++
		cnt = 1
	}
	if cnt > 1 {
		pos = compressItoa(chars, cnt, pos)
	}
	return pos
}

func compressItoa(chars []byte, cnt int, pos int) int {
	st := pos
	for cnt > 0 {
		chars[pos] = byte(cnt%10) + '0'
		pos++
		cnt /= 10
	}
	for en := pos - 1; st < en; st, en = st+1, en-1 {
		chars[st], chars[en] = chars[en], chars[st]
	}
	return pos
}

