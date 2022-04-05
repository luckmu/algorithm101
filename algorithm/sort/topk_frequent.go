package sort

// 347
// runtime 98.61%; memory 67.33%
func topKFrequent(nums []int, k int) []int {
	maxcnt := 0
	nr := make(map[int]int)
	for _, v := range nums {
		nr[v]++
		if nr[v] > maxcnt {
			maxcnt = nr[v]
		}
	}
	buckets := make([][]int, maxcnt+1)
	for k, v := range nr {
		buckets[v] = append(buckets[v], k)
	}
	rets := make([]int, 0)
	for i := maxcnt; i >= 0 && len(rets) < k; i-- {
		for _, v := range buckets[i] {
			rets = append(rets, v)
			if len(rets) >= k {
				break
			}
		}
	}
	return rets
}
