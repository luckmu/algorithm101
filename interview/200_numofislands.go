package interview

func draw(used [][]bool, grid [][]byte, x, y int) {
	used[y][x] = true
	if y != 0 && !used[y-1][x] && grid[y-1][x] == '1' {
		draw(used, grid, x, y-1)
	}
	if y != len(used)-1 && !used[y+1][x] && grid[y+1][x] == '1' {
		draw(used, grid, x, y+1)
	}
	if x != 0 && !used[y][x-1] && grid[y][x-1] == '1' {
		draw(used, grid, x-1, y)
	}
	if x != len(used[0])-1 && !used[y][x+1] && grid[y][x+1] == '1' {
		draw(used, grid, x+1, y)
	}
}

// time 86.66%; mem 34.67%
func numIslands(grid [][]byte) int {
	r := 0
	used := make([][]bool, len(grid))
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, len(grid[0]))
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if !used[y][x] && grid[y][x] == '1' {
				r++
				draw(used, grid, x, y)
			}
		}
	}
	return r
}

func Q200(grid [][]byte) int {
	return numIslands(grid)
}
