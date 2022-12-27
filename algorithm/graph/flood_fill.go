package graph

var _ = floodFill

func dfsFloodFill(i, j, startv, color int, image [][]int, traversed [][]bool) {
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
			image[newi][newj] == startv {
			traversed[newi][newj] = true
			image[newi][newj] = color
			dfsFloodFill(newi, newj, startv, color, image, traversed)
		}
	}
}

// 733
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	traversed := make([][]bool, len(image))
	for i := 0; i < len(traversed); i++ {
		traversed[i] = make([]bool, len(image[0]))
	}
	startv := image[sr][sc]
	image[sr][sc] = color
	traversed[sr][sc] = true
	dfsFloodFill(sr, sc, startv, color, image, traversed)
	return image
}
