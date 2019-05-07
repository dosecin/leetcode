/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */
func fizzBuzz(n int) []string {
	var ret []string
    for i := 1; i <= n; i++ {
		if i % 15 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if i % 5 == 0 {
			ret = append(ret, "Buzz")
		} else if i % 3 == 0 {
			ret = append(ret, "Fizz")
		} else {
			ret = append(ret, strconv.Itoa(i))
		}
	}
	return ret
}

