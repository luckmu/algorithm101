package greedy

import (
	"sort"
)

// 455
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var count int = 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] > s[j] {
			j++
		} else {
			count++
			i, j = i+1, j+1
		}
	}
	return count
}
