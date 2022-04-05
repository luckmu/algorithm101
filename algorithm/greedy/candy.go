package greedy

// 135
// 1. 每个孩子至少分配到 1 个糖果
// 2. 相邻两个孩子评分更高的孩子会获得更多的糖果
func candy(ratings []int) int {
	// init
	candies := make([]int, len(ratings))
	for i := 0; i < len(candies); i++ {
		candies[i] = 1
	}
	// left to right
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			candies[i+1]++
		}
	}
	// right to left
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candies[i-1]++
		}
	}
	// sum
	var ret int = 0
	for _, c := range candies {
		ret += c
	}

	return ret
}
