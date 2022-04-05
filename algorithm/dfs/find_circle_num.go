package dfs

// 547
func findCircleNum(isConnected [][]int) int {
	if len(isConnected) <= 0 || len(isConnected[0]) <= 0 {
		return 0
	}
	var circlenum int
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[0]); j++ {
			if isConnected[i][j] == 1 {
				// eliminate other cities in a province
				// 1. x connects to i
				// 2. j connects to x
				circlenum++
				clearprev(isConnected, i)
				clearnext(isConnected, j)
			}
		}
	}
	return circlenum
}

func clearprev(isConnected [][]int, dst int) {
	for i := 0; i < len(isConnected); i++ {
		// i -> dst
		if isConnected[i][dst] == 1 {
			isConnected[i][dst] = 0
			clearprev(isConnected, i)
		}
	}
}

func clearnext(isConnected [][]int, src int) {
	for j := 0; j < len(isConnected[0]); j++ {
		// src -> j
		if isConnected[src][j] == 1 {
			isConnected[src][j] = 0
			clearnext(isConnected, j)
		}
	}
}
