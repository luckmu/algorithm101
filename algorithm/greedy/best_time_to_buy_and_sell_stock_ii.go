package greedy

// 122
// runtime 100%; memory 80.18%
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var profit int
	for i := 0; i+1 < len(prices); {
		for j := i + 1; j < len(prices); {
			if prices[i] >= prices[j] {
				i, j = j, j+1
			} else {
				profit += prices[j] - prices[i]
				i = j
				break
			}
		}
	}
	return profit
}
