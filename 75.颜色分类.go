/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	/*基排序实现
	colors := make([]int, 3, 3)
	for _, n := range nums {
		colors[n]++
	}
	idx := 0
	for color, num := range colors {
		for num > 0 {
			nums[idx] = color
			idx++
			num--
		}
	}
	*/
	head := []int{-1, -1, -1}
	swap := func(i, j, iColor, jColor int) int {
		if head[jColor] == -1 {
			head[jColor] = i
		}
		nums[i], nums[j] = jColor, iColor
		head[iColor]++
		return head[iColor] - 1
	}
	for i, n := range nums {
		if n == 2 {
			if head[n] == -1 {
				head[n] = i
			}
			continue
		} else if n == 1 {
			if head[2] != -1 {
				swap(head[2], i, 2, n)
			} else if head[1] == -1 {
				head[1] = i
			}
		} else {
			if head[2] != -1 {
				i = swap(head[2], i, 2, n)
			}
			if head[1] != -1 {
				swap(head[1], i, 1, n)
			}
			if head[0] == -1 {
				head[0] = i
			}
		}
	}
}

// @lc code=end
