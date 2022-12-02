package test

import (
	"algorithm101/algorithm/sort"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	var nums = []int{1, 5, 3, 2, 6, 4, 457, 78}
	sort.Quicksort(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6, 78, 457}) {
		t.Error("quicksort has failed")
	}
	t.Log("nums:", nums)
}
