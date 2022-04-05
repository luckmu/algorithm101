package dfs

// 695
// runtime 67.76%; memory 82.06%
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	maxarea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				curarea := dfs(grid, i, j)
				if maxarea < curarea {
					maxarea = curarea
				}
			}
		}
	}
	return maxarea
}

var direction = []int{-1, 0, 1, 0, -1}

func dfs(grid [][]int, r, c int) int {
	// 递归的部分，是一定要做判断的；
	// 不光是main里面做了判断，自身调用会改变条件s
	if grid[r][c] == 0 {
		return 0
	}
	area := grid[r][c] // var area int = 1
	// why? fuck, marking grid[r][c] was used
	grid[r][c] = 0
	for i := 0; i < 4; i++ {
		x, y := r+direction[i], c+direction[i+1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			area += dfs(grid, x, y)
		}
	}
	return area
}
