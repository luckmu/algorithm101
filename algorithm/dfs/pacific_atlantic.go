package dfs

// 417
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) <= 0 || len(heights[0]) <= 0 {
		return [][]int{}
	}
	used := make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		used[i] = make([]bool, len(heights[0]))
	}
	var res = make([][]int, 0)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			used[i][j] = true
			if pacific(heights, used, i, j) && atlantic(heights, used, i, j) {
				res = append(res, []int{i, j})
			}
			used[i][j] = false
		}
	}
	return res
}

func pacific(heights [][]int, used [][]bool, i, j int) bool {
	if i == 0 || j == 0 {
		// left&top
		return true
	}
	var can bool
	for k := 0; k < 4; k++ {
		x, y := i+direction[k], j+direction[k+1]
		if x >= 0 && x < len(heights) && y >= 0 && y < len(heights[0]) && !used[x][y] {
			if heights[x][y] <= heights[i][j] {
				// can flow to lower place x,y
				used[x][y] = true
				can = can || pacific(heights, used, x, y)
				used[x][y] = false
			}
		}
	}
	return can
}

func atlantic(heights [][]int, used [][]bool, i, j int) bool {
	if i == len(heights)-1 || j == len(heights[0])-1 {
		// right&bottom
		return true
	}
	var can bool
	for k := 0; k < 4; k++ {
		x, y := i+direction[k], j+direction[k+1]
		if x >= 0 && x < len(heights) && y >= 0 && y < len(heights[0]) && !used[x][y] {
			if heights[x][y] <= heights[i][j] {
				used[x][y] = true
				can = can || atlantic(heights, used, x, y)
				used[x][y] = false
			}
		}
	}
	return can
}
