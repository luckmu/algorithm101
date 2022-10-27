package interview

// math 类的题目, 主要是推导 overflow 的条件, 可跳过此题目
func reverse(x int) int {
	minint32, maxint32 := -1<<31, 1<<31-1
	// 2147483647 -2147483648
	r := 0
	for x != 0 {
		num := x % 10
		x /= 10
		if r > maxint32/10 || r == maxint32/10 && num > 7 {
			// 10r + num > max
			// 10r + num > 2147483640 + 7, num->[0,9]
			// r > max/10 || r == max/10 && num > 7
			return 0
		}
		if r < minint32/10 || r == minint32/10 && num < -8 {
			// 10r + num < min
			// 10r + num < -2147483640 - 8, num->[0,9]
			// r < min/10 || r == min/10 && num < -8
			return 0
		}
		r = 10*r + num
	}
	return r
}

func Q7(x int) int {
	// time 100.00%
	// mem  86.53%
	return reverse(x)
}
