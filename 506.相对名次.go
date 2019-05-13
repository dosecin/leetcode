/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */
var rankstr = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(nums []int) []string {
	cp := append([]int{}, nums...)
	sort.Ints(cp)
	cpm := map[int]string{}
	for i := len(cp) - 1; i >= 0; i-- {
		rank := len(cp) - 1 - i
		if rank < 3 {
			cpm[cp[i]] = rankstr[rank]
		} else {
			cpm[cp[i]] = strconv.Itoa(rank + 1)
		}
	}
	res := []string{}
	for _, n := range nums {
		res = append(res, cpm[n])
	}
	return res
}

