package interview

// double pointer + greedy
func maxArea(height []int) int {
	area := 0
	for l, r := 0, len(height)-1; l < r; {
		tmp := r - l
		if height[l] < height[r] {
			tmp *= height[l]
			l++
		} else {
			tmp *= height[r]
			r--
		}
		if tmp > area {
			area = tmp
		}
	}
	return area
}

func Q11(height []int) int {
	// time 26.79%
	// mem  85.74%
	return maxArea(height)
}
