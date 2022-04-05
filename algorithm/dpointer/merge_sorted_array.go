package dpointer

// 88

// runtime 100%, memory 83.83%
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}
	if n == 0 {
		return
	}
	// from end to start, avoid overwriting
	var i, j int = m - 1, n - 1
	for k := m + n - 1; k >= 0; k-- {
		if i < 0 {
			copy(nums1[:k+1], nums2[:j+1])
			return
		} else if j < 0 {
			copy(nums1[:k+1], nums1[:i+1])
			return
		} else {
			if nums1[i] > nums2[j] {
				nums1[k], i = nums1[i], i-1
			} else {
				nums1[k], j = nums2[j], j-1
			}
		}
	}
}
