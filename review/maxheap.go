package review

var (
	_ = heapsort
	_ = maxheapify
)

func heapsort(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		maxheapify(nums, i)
	}
}

func maxheapify(nums []int, i int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < len(nums) && nums[left] > nums[largest] {
		largest = left
	}
	if right < len(nums) && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		maxheapify(nums, largest)
	}
}
