package sort

func quicksort(nums []int) {
	// quicksorti(nums, 0, len(nums)-1)
	quicksortii(nums, 0, len(nums)-1)
}

func quicksorti(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	pivot := nums[i]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quicksorti(nums, l, i-1)
	quicksorti(nums, i+1, r)
}

// quicksortii method2
// 双指针i，j，从index 0开始，大于pivot的全放后面（实际操作，小于pivot的全放前面）
// 1. pivot := nums[r]
// 2. nums[j] < nums[i] 则 swap(i, j);i++; 否则 j++
// 3. 当 j==r 时, swap(i, r), nums[i] == pivot
// 4. partition(nums, i, i-1); partition(nums, i+1, r)
func quicksortii(nums []int, l, r int) {
	// 1. l == 0 && r == -1
	// 2. l == 1 && r == 0
	if l >= r {
		return
	}
	i, pivot := l, nums[r]
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	quicksortii(nums, l, i-1)
	quicksortii(nums, i+1, r)
}
