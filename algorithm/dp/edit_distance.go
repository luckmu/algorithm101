package dp

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word1)+1)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				// replace
				if word2[i-1] == word1[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				// add
				if dp[i-1][j]+1 < dp[i][j] {
					dp[i][j] = dp[i-1][j] + 1
				}
				if dp[i][j-1]+1 < dp[i][j] {
					dp[i][j] = dp[i][j-1] + 1
				}
			}
		}
	}
	return dp[len(word2)][len(word1)]
}
