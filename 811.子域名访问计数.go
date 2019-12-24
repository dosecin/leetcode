import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	tmpAns := map[string]int{}
	for _, domain := range cpdomains {
		dom := strings.Split(domain, " ")
		cnt, _ := strconv.Atoi(dom[0])
		tmpAns[dom[1]] += cnt
		for {
			dom = strings.SplitN(dom[1], ".", 2)
			if len(dom) < 2 {
				break
			}
			tmpAns[dom[1]] += cnt
		}
	}

	ans := []string{}
	for d, c := range tmpAns {
		ans = append(ans, fmt.Sprintf("%d %s", c, d))
	}
	return ans
}

// @lc code=end
