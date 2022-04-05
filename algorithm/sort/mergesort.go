package sort

func mergesort(nums []int) {
	mergemain(nums, 0, len(nums)-1)
}

func mergemain(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergemain(nums, l, mid)   // []
	mergemain(nums, mid+1, r) // (]
	tmp := make([]int, r-l+1)
	i, j := l, mid+1
	for i <= mid && j <= r {
		tidx := i - l + j - mid - 1
		if nums[i] <= nums[j] {
			tmp[tidx] = nums[i]
			i++
		} else {
			tmp[tidx] = nums[j]
			j++
		}
	}
	if i <= mid {
		copy(tmp[i-l+j-mid-1:], nums[i:mid+1])
	}
	if j <= r {
		copy(tmp[i-l+j-mid-1:], nums[j:r+1])
	}
	copy(nums[l:r+1], tmp)
}
