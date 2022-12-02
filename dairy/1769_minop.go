package dairy

func addzero(ans *[]int, boxes string, i int) {
	// add zero, only add ans[i]
	*ans = append(*ans, 0)
	for j := i - 1; j >= 0; j-- {
		if boxes[j] == '1' {
			(*ans)[i] += (i - j)
		}
	}
}

func addone(ans *[]int, boxes string, i int) {
	for j := 0; j < i; j++ {
		(*ans)[j] += i - j
	}
	addzero(ans, boxes, i)
}

func minOperations(boxes string) []int {
	ans := []int{}
	for i := 0; i < len(boxes); i++ {
		if boxes[i] == '0' {
			addzero(&ans, boxes, i)
		}
		if boxes[i] == '1' {
			addone(&ans, boxes, i)
		}
	}
	return ans
}
