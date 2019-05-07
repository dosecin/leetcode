/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

var (
	hm = map[int][]string{}
	mm = map[int][]string{}
)

func init() {
	for i := 0; i < 12; i++ {
		n, cnt := i, 0
		for n != 0 {
			cnt++
			n = n & (n - 1)
		}
		hm[cnt] = append(hm[cnt], fmt.Sprintf("%d", i))
	}

	for i := 0; i < 60; i++ {
		n, cnt := i, 0
		for n != 0 {
			cnt++
			n = n & (n - 1)
		}
		mm[cnt] = append(mm[cnt], fmt.Sprintf("%02d", i))
	}
}

func readBinaryWatch(num int) []string {
	h, m := num, 0
	var ret []string
	for {
		if h < 0 {
			break
		}
		if h < 4 && m < 6 {
			hv, mv := hm[h], mm[m]
			for _, hs := range hv {
				for _, ms := range mv {
					ret = append(ret, fmt.Sprintf("%s:%s", hs, ms))
				}
			}
		}
		h--
		m++
	}
	return ret
}

