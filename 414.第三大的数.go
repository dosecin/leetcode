/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	maxNums := []int{}
	for _, num := range nums {
		if len(maxNums) < 3 {
			insertIdx := 0
			for _, mn := range maxNums {
				if num > mn {
					break
				} else if num == mn {
					insertIdx = -1
					break
				}
				insertIdx++
			}
			if insertIdx != -1 {
				if insertIdx == len(maxNums) {
					maxNums = append(maxNums, num)
				} else {
					maxNums = append(maxNums, num)
					copy(maxNums[insertIdx+1:], maxNums[insertIdx:])
					maxNums[insertIdx] = num
				}
			}
			continue
		}
		for i := 0; i < len(maxNums); i++ {
			if num > maxNums[i] {
				copy(maxNums[i+1:], maxNums[i:])
				maxNums[i] = num
				break
			} else if num == maxNums[i] {
				break
			}
		}
	}
	if len(maxNums) == 3 {
		return maxNums[2]
	}
	return maxNums[0]
}

// @lc code=end
