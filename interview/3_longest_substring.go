package interview

func slide(s string, idx int) int {
	m := map[byte]int{}
	for i := idx; i >= 0; i-- {
		if m[s[i]] >= 1 {
			// repeated
			return idx - i
		}
		m[s[i]]++ // 'b':0 -> 'b':1
	}
	return idx + 1
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if r := slide(s, i); r > dp[i-1] {
			dp[i] = r
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)-1]
}

func Q3(s string) int {
	return lengthOfLongestSubstring(s)
}
