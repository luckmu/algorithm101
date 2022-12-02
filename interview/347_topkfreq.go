package interview

import (
	"math/rand"
	"time"
)

func mqsort(m map[int]int, nums []int, start, end int) {
	if start >= end {
		return
	}
	rand.Seed(time.Hour.Nanoseconds())
	pos := start + rand.Intn(end-start+1) // start + offset = pivot position
	nums[start], nums[pos] = nums[pos], nums[start]
	pivot := nums[start]
	l, r := start, end
	for l < r {
		for l < r && m[nums[r]] >= m[pivot] {
			r--
		}
		nums[l] = nums[r]
		for l < r && m[nums[l]] <= m[pivot] {
			l++
		}
		nums[r] = nums[l]
	}
	// l == r
	nums[l] = pivot
	mqsort(m, nums, start, l-1)
	mqsort(m, nums, l+1, end)
}

// exceed time limit
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	// qsort
	mqsort(m, nums, 0, len(nums)-1)
	ans := []int{nums[len(nums)-1]}
	delete(m, ans[0])
	for i := len(nums) - 2; i >= 0 && len(ans) < k; i-- {
		if _, ok := m[nums[i]]; ok {
			ans = append(ans, nums[i])
			delete(m, nums[i])
		}
	}
	return ans
}

func Q347(nums []int, k int) []int {
	if time.Now().Second()%2 == 0 {
		return topKFrequent(nums, k)
	}
	return topKFrequentii(nums, k)
}

// time 7.55%; mem 84.91%
func topKFrequentii(nums []int, k int) []int {
	freq := map[int]int{}
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}
	heapSize := len(nums)
	buildMaxHeapii(nums, heapSize, freq)
	used := map[int]struct{}{}
	ans := []int{}
	for i := len(nums) - 1; i >= 0 && len(ans) < k; i-- {
		if _, ok := used[nums[0]]; !ok {
			ans = append(ans, nums[0])
			used[nums[0]] = struct{}{}
		}
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapifyii(nums, 0, heapSize, freq)
	}
	return ans
}

func buildMaxHeapii(nums []int, heapSize int, freq map[int]int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapifyii(nums, i, heapSize, freq)
	}
}

func maxHeapifyii(nums []int, i, heapSize int, freq map[int]int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < heapSize && freq[nums[left]] > freq[nums[largest]] {
		largest = left
	}
	if right < heapSize && freq[nums[right]] > freq[nums[largest]] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapifyii(nums, largest, heapSize, freq)
	}
}
