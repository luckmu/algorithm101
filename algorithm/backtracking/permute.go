package backtracking

func permute(nums []int) [][]int {
	ret, used := make([][]int, 0), make([]bool, len(nums))
	backtracking(nums, []int{}, used, &ret)
	return ret
}

func backtracking(nums []int, cur []int, used []bool, res *[][]int) {
	if len(cur) >= len(nums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			backtracking(nums, append(cur, nums[i]), used, res)
			used[i] = false
		}
	}
}
