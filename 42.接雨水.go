/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	maxIdx := 0
	for i, h := range height {
		if h > height[maxIdx] {
			maxIdx = i
		}
	}
	calcArea := func(left, right int) int {
		w := right - left - 1
		h := height[left]
		if h > height[right] {
			h = height[right]
		}
		area := w * h
		left++
		for left < right {
			area -= height[left]
			left++
		}
		return area
	}
	sumArea := 0
	l, r := 0, 1
	for r <= maxIdx {
		if height[r] >= height[l] {
			sumArea += calcArea(l, r)
			l = r
		}
		r++
	}
	l, r = len(height)-2, len(height)-1
	for l >= maxIdx {
		if height[r] <= height[l] {
			sumArea += calcArea(l, r)
			r = l
		}
		l--
	}
	return sumArea
}

// @lc code=end
