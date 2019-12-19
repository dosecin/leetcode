/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	numArr := []int{}
	r, pre := num, 0
	max, maxIdx := 0, -1
	inv, invIdx := 0, -1
	for r != 0 {
		b := r % 10
		if b > pre {
			max, maxIdx = b, len(numArr)
		} else if b < pre && max > inv {
			inv, invIdx = max, maxIdx
		}
		numArr = append(numArr, b)
		pre = b
		r /= 10
	}
	if invIdx != -1 {
		r := 0
		for i := len(numArr) - 1; i >= 0; i-- {
			if numArr[i] < inv {
				numArr[i], numArr[invIdx] = numArr[invIdx], numArr[i]
				for j := i; j >= 0; j-- {
					r = r*10 + numArr[j]
				}
				break
			}
			r = r*10 + numArr[i]
		}
		return r
	}
	return num
}

// @lc code=end
