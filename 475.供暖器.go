/*
 * @lc app=leetcode.cn id=475 lang=golang
 *
 * [475] 供暖器
 */
func findRadius(houses []int, heaters []int) int {
	if len(houses) == 0 || len(heaters) == 0 {
		return 0
	}
	sort.Ints(houses)
	sort.Ints(heaters)
	radius := heaters[0] - houses[0]
	l, r := 0, 0
	for i := range houses {
		for houses[i] > heaters[r] {
			l = r
			r++
			if r >= len(heaters) {
				r--
				break
			}
		}
		nl := absInt(heaters[l] - houses[i])
		nr := absInt(heaters[r] - houses[i])
		minR := nl
		if minR > nr {
			minR = nr
		}
		if minR > radius {
			radius = minR
		}
	}
	return radius
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

