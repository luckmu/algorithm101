package interview

// time 41.67%; mem 12.78%
// func maxSubArrayii(nums []int) int {
// 	dp, r := make([]int, len(nums)), nums[0]
// 	dp[0] = nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		// 状态转移方程？why
// 		// 这里的 dp 起什么作用？
// 		// 并不是传统的 `return dp[len(dp)-1]`
// 		if dp[i-1] > 0 {
// 			dp[i] = dp[i-1] + nums[i]
// 		} else {
// 			dp[i] = nums[i]
// 		}
// 		if r < dp[i] {
// 			r = dp[i]
// 		}
// 	}
// 	return r
// }

// time 35.61%; mem 48.30%
func maxSubArray(nums []int) int {
	var max int = nums[0]
	var now int
	for i := 0; i < len(nums); i++ {
		now += nums[i]
		if now > max {
			max = now
		}
		if now < 0 {
			now = 0
		}
	}
	return max
}

func Q53(nums []int) int {
	return maxSubArray(nums)
}
