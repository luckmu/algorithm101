package tree

import "fmt"

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}
	ans := make([]float64, 0)
	qs := []*TreeNode{root}
	for len(qs) > 0 {
		var lvlcount, lvlsum int = len(qs), 0
		for i := 0; i < lvlcount; i++ {
			n := qs[0]
			qs = qs[1:]
			fmt.Println(lvlsum, n.Val)
			lvlsum += n.Val
			if n.Left != nil {
				qs = append(qs, n.Left)
			}
			if n.Right != nil {
				qs = append(qs, n.Right)
			}
		}
		ans = append(ans, float64(lvlsum)/float64(lvlcount))
	}
	return ans
}
