import "sort"

/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	taskCnt := make([]int, 26)
	for _, t := range tasks {
		taskCnt[t-'A']++
	}
	sort.Ints(taskCnt)
	idleCnt := (taskCnt[25] - 1) * n
	time := idleCnt + taskCnt[25]
	for i := len(taskCnt) - 2; i >= 0; i-- {
		if taskCnt[i] == taskCnt[25] {
			time++
			idleCnt -= taskCnt[25] - 1
		} else {
			idleCnt -= taskCnt[i]
		}
		if idleCnt <= 0 {
			return len(tasks)
		}
	}
	return time
}

// @lc code=end
