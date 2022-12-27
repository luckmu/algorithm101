package graph

var _ = numIslands

func dfs(i, j int, traversed [][]bool, grid [][]byte) {
	var newi, newj int
	for k := 0; k < 4; k++ {
		switch k {
		case 0:
			newi, newj = i+1, j
		case 1:
			newi, newj = i-1, j
		case 2:
			newi, newj = i, j+1
		case 3:
			newi, newj = i, j-1
		}
		if newi >= 0 && newi < len(traversed) &&
			newj >= 0 && newj < len(traversed[0]) &&
			!traversed[newi][newj] &&
			grid[newi][newj] == '1' {
			traversed[newi][newj] = true
			dfs(newi, newj, traversed, grid)
		}
	}
}

func numIslands(grid [][]byte) int {
	ans := 0
	traversed := make([][]bool, len(grid))
	for i := 0; i < len(traversed); i++ {
		traversed[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !traversed[i][j] && grid[i][j] == '1' {
				// 点 (i, j) 没有被遍历过 && 是陆地
				ans++                      // 结果+1
				traversed[i][j] = true     // 当前点遍历
				dfs(i, j, traversed, grid) // 排除所有连通点 -> traversed[i][j] = true
			}
		}
	}
	return ans
}
