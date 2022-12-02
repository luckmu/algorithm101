package interview

// time 100%; mem 61.95%
func findpeak(start, end int, nums []int) int {
	if start > end {
		return -1
	}
	i := (start + end) / 2
	if (i <= 0 || nums[i-1] < nums[i]) &&
		(i >= len(nums)-1 || nums[i] > nums[i+1]) {
		return i
	}
	if i := findpeak(start, i-1, nums); i != -1 {
		return i
	}
	if i := findpeak(i+1, end, nums); i != -1 {
		return i
	}
	return -1
}

// log(n) 2分
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	return findpeak(0, len(nums)-1, nums)
}

func Q162(nums []int) int {
	return findPeakElement(nums)
}
