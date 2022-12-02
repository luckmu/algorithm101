package interview

// time 85.81%; mem 93.89%
func longestConsecutive(nums []int) int {
	have := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		have[nums[i]] = struct{}{}
	}
	maxl := 0
	for num := range have {
		curl := 1
		delete(have, num)
		for r := 1; ; r++ {
			if _, ok := have[num+r]; !ok {
				break
			}
			curl++
			delete(have, num+r)
		}
		for l := 1; ; l++ {
			if _, ok := have[num-l]; !ok {
				break
			}
			curl++
			delete(have, num-l)
		}
		if curl > maxl {
			maxl = curl
		}
	}
	return maxl
}

func Q128(nums []int) int {
	return longestConsecutive(nums)
}
