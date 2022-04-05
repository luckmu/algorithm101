package bsearch

// 69
// runtime 100%; memory 99.96%
func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	if x <= 3 {
		return 1
	}
	l, r := 0, x
	// mid cannot equal to 0
	for l < r {
		mid := l + (r-l)/2
		if mid == l {
			return mid
		}
		k := x / mid
		if k == mid {
			return mid
		} else if k > mid {
			l = mid
		} else {
			r = mid
		}
	}
	return 0
}
