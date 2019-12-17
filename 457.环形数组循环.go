/*
 * @lc app=leetcode.cn id=457 lang=golang
 *
 * [457] 环形数组循环
 */

// @lc code=start
func circularArrayLoop(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	calcIndex := func(idx int) int {
		return ((idx+nums[idx])%len(nums) + len(nums)) % len(nums)
	}

	setZero := func(i int) {
		for {
			last := i
			i = calcIndex(i)
			if nums[i] == 0 || nums[last]*nums[i] < 0 {
				nums[last] = 0
				break
			}
			nums[last] = 0
		}
	}

	for i := range nums {
		fast, slow := i, i
		for {
			last := slow
			slow = calcIndex(slow)
			if last == slow || nums[slow] == 0 || nums[last]*nums[slow] < 0 {
				setZero(i)
				break
			}

			last = fast
			fast = calcIndex(fast)
			if last == fast || nums[fast] == 0 || nums[last]*nums[fast] < 0 {
				setZero(i)
				break
			}

			last = fast
			fast = calcIndex(fast)
			if last == fast || nums[fast] == 0 || nums[last]*nums[fast] < 0 {
				setZero(i)
				break
			}

			if fast == slow {
				return true
			}
		}
	}
	return false
}

// @lc code=end
