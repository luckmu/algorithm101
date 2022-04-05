package backtracking

// 77
func combine(n int, k int) [][]int {
	var ret = make([][]int, 0)
	backtrackingii(&ret, []int{}, 1, n, k)
	return ret
}

func backtrackingii(ret *[][]int, cur []int, pos, n, k int) {
	if len(cur) == k {
		var tmp = make([]int, len(cur))
		copy(tmp, cur)
		*ret = append(*ret, tmp)
		return
	}
	for i := pos; i <= n; i++ {
		// i 已经使用了，下个递归应该从用 i+1 开始
		backtrackingii(ret, append(cur, i), i+1, n, k)
	}
}
