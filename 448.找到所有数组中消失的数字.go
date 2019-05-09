/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */
func findDisappearedNumbers(nums []int) []int {
    for i := range nums {
		v := nums[i]
		if v == 0 || v == nums[v-1] {
			if i != v-1 {
				nums[i] = 0
			}
			continue
		}
		for {
			if v == 0 {
				break
			}
			if v == nums[v-1] {
				if i != v-1 {
					nums[i] = 0
				}
				break
			}
			nums[i], nums[v-1] = nums[v-1], nums[i]
			v = nums[i]
		}
	}
	res := []int{}
	for i := range nums {
		if nums[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}

