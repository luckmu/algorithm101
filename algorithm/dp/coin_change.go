package dp

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	// when amount is 0; need 0 coins;
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			} else if coin == i {
				dp[i] = 1
			} else {
				// coin < i
				if dp[i-coin] == -1 {
					continue
				} else {
					// dp[i-coin] != -1
					if dp[i] == -1 {
						dp[i] = dp[i-coin] + 1
					} else if dp[i] > dp[i-coin]+1 {
						dp[i] = dp[i-coin] + 1
					}
				}
			}
		}
	}
	return dp[amount]
}
