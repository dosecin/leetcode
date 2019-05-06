/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */
func reverseVowels(s string) string {
    y := map[byte]struct{}{
		'a':struct{}{},
		'e':struct{}{},
		'i':struct{}{},
		'o':struct{}{},
		'u':struct{}{},
		'A':struct{}{},
		'E':struct{}{},
		'I':struct{}{},
		'O':struct{}{},
		'U':struct{}{},
	}
	bs := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j {
			_, ok := y[bs[i]]
			if ok {
				break
			}
			i++
		}
		for i < j {
			_, ok := y[bs[j]]
			if ok {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		bs[i],bs[j]=bs[j],bs[i]
		i++
		j--
	}
	return string(bs)
}

