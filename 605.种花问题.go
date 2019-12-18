/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	pos := 0
	for pos < len(flowerbed) {
		if flowerbed[pos] == 0 {
			if pos+1 >= len(flowerbed) || flowerbed[pos+1] == 0 {
				n--
				if n == 0 {
					return true
				}
				pos += 2
			} else {
				pos += 3
			}
		} else {
			pos += 2
		}
	}
	return false
}

// @lc code=end
