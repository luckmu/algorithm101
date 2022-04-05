package dpointer

import (
	"sort"
)

// 122

// runtime 94.15%, memory 58.78%
func twoSum(numbers []int, target int) []int {
	sort.Ints(numbers)
	i, j := 0, len(numbers)-1
	for i < j {
		if sum := numbers[i] + numbers[j]; sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}
