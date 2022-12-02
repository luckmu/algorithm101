package interview

// time 25.13%; mem 99.92%
func maxProduct(nums []int) int {
	// dpmax[i] = max(nums[i], dpmax[i-1]*nums[i], dpmin[i-1]*nums[i])
	// dpmin[i] = min(nums[i], dpmax[i-1]*nums[i], dpmin[i-1]*nums[i])
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Q152(nums []int) int {
	return maxProduct(nums)
}
