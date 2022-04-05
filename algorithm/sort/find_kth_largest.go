package sort

// 215
func findKthLargest(nums []int, k int) int {
	quicksort(nums)
	return nums[len(nums)-k]
}
