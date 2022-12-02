package interview

// time 100%; mem 92.13%
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	p, pp := nums[1], nums[0]
	if pp > p {
		p = pp
	}
	for i := 2; i < len(nums); i++ {
		tmp := 0
		if pp+nums[i] > p {
			tmp = pp + nums[i]
		} else {
			tmp = p
		}
		p, pp = tmp, p
	}
	return p
}

func Q198(nums []int) int {
	return rob(nums)
}
