package interview

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		// 从二叉树最底部往上
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	// nums[i] = nums[largest]
	// nums[i] 为 "根, 左, 右" 中的最大值
	if largest != i {
		// 右、左子节点和父节点交换
		nums[i], nums[largest] = nums[largest], nums[i]
		// 当前节点为largest, 交换下去节点构成的子树也要保证最大堆
		maxHeapify(nums, largest, heapSize)
	}
}

// [0,1,2,3]
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func Q215(nums []int, k int) int {
	return findKthLargest(nums, k)
}
