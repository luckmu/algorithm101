package backtracking

var Q = "Q"
var DOT = "."

// 51
// just follow the standart solution
func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	ret := make([][]string, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

		}
	}

	return ret
}

func backtrackingiiii(ret *[][]string, board []string, column, ldiag, rdiag []bool, row, n int) {
	// every pos, put queen or not
	if row == n {
		tmp := make([]string, len(board))
		copy(tmp, board)
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if column[i] || ldiag[n-row+1] || rdiag[row+1] {
			continue
		}

		// board[row][i] = 'Q'
		column[i] = true
		ldiag[n-row+i-1] = true
		rdiag[row+i] = true
		backtrackingiiii()
	}

}
