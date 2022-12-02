package interview

// certainly time limit exceeded
// func canJumpdfs(i int, nums []int) bool {
// 	if i+nums[i] >= len(nums)-1 {
// 		return true
// 	}
// 	var can bool
// 	for offset := 0; offset < nums[i]; offset++ {
// 		can = can || canJumpdfs(i+offset, nums)
// 	}
// 	return can
// }

// time 81.20%; mem 93.59%
func canJump(nums []int) bool {
	if len(nums) == 1 && nums[0] == 0 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		if dp[i] {
			if i+nums[i] >= len(nums)-1 {
				return true
			}
			for j := 1; j <= nums[i]; j++ {
				dp[i+j] = true
			}
		}
	}
	return false
}

func Q55(nums []int) bool {
	// return canJumpdfs(0, nums)
	return canJump(nums)
}
