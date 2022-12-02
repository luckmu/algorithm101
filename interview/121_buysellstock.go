package interview

// O(n^2), excceed time limit
// func maxProfitOn2(prices []int) int {
// 	// -> max difference between 2 idx
// 	if len(prices) <= 1 {
// 		return 0
// 	}
// 	maxp := 0
// 	for i := 0; i < len(prices)-1; i++ {
// 		if prices[i+1] <= prices[i] {
// 			continue
// 		}
// 		for j := i + 1; j < len(prices); j++ {
// 			if prices[j]-prices[i] > maxp {
// 				maxp = prices[j] - prices[i]
// 			}
// 		}
// 	}
// 	return maxp
// }

// time 47.80%; mem 35.49%
func maxProfit(prices []int) int {
	// -> max difference between 2 idx
	if len(prices) <= 1 {
		return 0
	}
	// dp
	// check maxp && update mini
	maxp, mini := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[mini] > maxp {
			maxp = prices[i] - prices[mini]
		}
		if prices[i] < prices[mini] {
			mini = i
		}
	}
	return maxp
}

func Q121(prices []int) int {
	return maxProfit(prices)
}
