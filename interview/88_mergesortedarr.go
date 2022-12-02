package interview

// time 15.75%; mem 35.19%
func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for k := m + 1 + n + 1 - 1; k >= 0; k-- {
		if m < 0 {
			copy(nums1, nums2[:n+1])
			return
		}
		if n < 0 {
			return
		}
		if nums1[m] < nums2[n] {
			nums1[k] = nums2[n]
			n--
		} else {
			nums1[k] = nums1[m]
			m--
		}
	}
}

func Q88(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}
