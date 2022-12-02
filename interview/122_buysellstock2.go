package interview

// time 85.79%; mem 99.95%
func maxProfit2(prices []int) int {
	// buy  at i
	// sell at j (j > i)
	// ->
	// buy  at i
	// sell at j && buy at j: if prices[j+n]>prices[j], sell at j+n else sell at j
	// every step i, we buy, if prices[j] > prices[i], we sell
	maxp := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			maxp += prices[i+1] - prices[i]
		}
	}
	return maxp
}

func Q122(prices []int) int {
	return maxProfit2(prices)
}
