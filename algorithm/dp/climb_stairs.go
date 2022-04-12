package dp

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	var cur, pre2, pre1 int = 0, 1, 2
	for i := 2; i < n; i++ {
		cur = pre1 + pre2
		pre1, pre2 = cur, pre1
	}
	return cur
}
