package backtracking

var direction = []int{-1, 0, 1, 0, -1}

// 79
func exist(board [][]byte, word string) bool {
	if len(word) <= 0 {
		return true
	}
	if len(board) <= 0 || len(board[0]) <= 0 {
		return false
	}
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				used[i][j] = true
				if backtrackingiii(board, used, word, i, j, 1) {
					return true
				}
				used[i][j] = false
			}
		}
	}
	return false
}

func backtrackingiii(board [][]byte, used [][]bool, word string, i, j, wordcnt int) bool {
	if wordcnt == len(word) {
		return true
	}
	for k := 0; k < 4; k++ {
		x, y := i+direction[k], j+direction[k+1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && !used[x][y] {
			if board[x][y] == word[wordcnt] {
				used[x][y] = true
				if backtrackingiii(board, used, word, x, y, wordcnt+1) {
					return true
				}
				used[x][y] = false
			}
		}
	}
	return false
}
