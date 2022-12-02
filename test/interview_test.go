package test

import (
	"algorithm101/interview"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type Node = interview.ListNode

func TestInterview(t *testing.T) {
	// // test: q3
	// m3 := map[string]int{
	// 	"abcabcbb": 3,
	// 	"bbbbb":    1,
	// 	"pwwkew":   3,
	// 	"au":       2,
	// }
	// for input, want := range m3 {
	// 	get := interview.Q3(input)
	// 	if want != get {
	// 		t.Fatal("input:", input, "want:", want, "but get:", get)
	// 	}
	// }

	// // test: q5
	// m5 := map[string]string{
	// 	"ac":      "a",
	// 	"babad":   "bab",
	// 	"cbbbd":   "bbb",
	// 	"bb":      "bb",
	// 	"bananas": "anana",
	// }
	// for input, want := range m5 {
	// 	get := interview.Q5(input)
	// 	if want != get {
	// 		t.Fatal("input:", input, "want:", want, "but get:", get)
	// 	}
	// }

	// // test: q7
	// m7 := map[int]int{
	// 	123:  321,
	// 	-123: -321,
	// 	120:  21,
	// }
	// for input, want := range m7 {
	// 	get := interview.Q7(input)
	// 	if want != get {
	// 		t.Fatal("input:", input, "want:", want, "but get:", get)
	// 	}
	// }

	// // test: q11
	// i11 := [][]int{{1, 8, 6, 2, 5, 4, 8, 3, 7}, {1, 1}}
	// r11 := []int{49, 1}
	// for i := 0; i < len(i11); i++ {
	// 	input, want := i11[i], r11[i]
	// 	if get := interview.Q11(input); get != want {
	// 		t.Fatal("input:", input, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q15
	// i15 := [][]int{{-1, 0, 1, 2, -1, -4}, {0, 1, 1}, {0, 0, 0}, {0, 0, 0, 0}}
	// r15 := [][][]int{{{-1, -1, 2}, {-1, 0, 1}}, {}, {{0, 0, 0}}, {{0, 0, 0}}}
	// for i, input := range i15 {
	// 	// 要求了顺序, 方便直接用 DeepEqual
	// 	if get, want := interview.Q15(input), r15[i]; !reflect.DeepEqual(get, want) {
	// 		t.Fatal("input:", input, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q28
	// m28 := map[[2]string]int{
	// 	{"sadbutsad", "sad"}:       0,
	// 	{"leetcode", "leeto"}:      -1,
	// 	{"aaa", "aaaa"}:            -1,
	// 	{"aaabaaabbbabaa", "babb"}: -1,
	// 	{"mississippi", "issip"}:   4,
	// }
	// for input, want := range m28 {
	// 	haystack, needle := input[0], input[1]
	// 	if get := interview.Q28(haystack, needle); get != want {
	// 		t.Fatal("input:", haystack, needle, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q33
	// i330 := [][]int{{4, 5, 6, 7, 0, 1, 2}, {4, 5, 6, 7, 0, 1, 2}, {1}, {1, 3}}
	// i331 := []int{0, 3, 1, 2}
	// o33 := []int{4, -1, 0, -1}
	// for i := 0; i < len(o33); i++ {
	// 	nums, target, want := i330[i], i331[i], o33[i]
	// 	if get := interview.Q33(nums, target); get != want {
	// 		t.Fatal("input:", nums, target, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q34
	// i34 := [][]int{{8, 5, 7, 7, 8, 8, 10}, {6, 5, 7, 7, 8, 8, 10}, {0}, {1, 1}, {2, 2, 2}}
	// o34 := [][]int{{3, 4}, {-1, -1}, {-1, -1}, {0, 0}, {0, 1}}
	// for i := 0; i < len(o34); i++ {
	// 	target, nums, want := i34[i][0], i34[i][1:], o34[i]
	// 	if get := interview.Q34(nums, target); !reflect.DeepEqual(want, get) {
	// 		t.Fatal("input:", nums, target, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q53
	// m53 := [][]int{{6, -2, 1, -3, 4, -1, 2, 1, -5, 4},
	// 	{1, 1}, {23, 5, 4, -1, 7, 8}, {-1, -2, -1},
	// 	{21, 8, -19, 5, -4, 20}}
	// for _, m := range m53 {
	// 	want, nums := m[0], m[1:]
	// 	if get := interview.Q53(nums); get != want {
	// 		t.Fatal("input:", nums, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q55
	// i55 := [][]int{{2, 3, 1, 1, 4}, {3, 2, 1, 0, 4}, {0}, {1}}
	// w55 := []bool{true, false, true, true}
	// for i := 0; i < len(w55); i++ {
	// 	nums, want := i55[i], w55[i]
	// 	if get := interview.Q55(nums); want != get {
	// 		t.Fatal("input:", nums, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q70
	// m70 := map[int]int{
	// 	2: 2,
	// 	3: 3,
	// }
	// for n, want := range m70 {
	// 	if get := interview.Q70(n); get != want {
	// 		t.Fatal("input:", n, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q76
	// m76 := map[[2]string]string{
	// 	{"ADOBECODEBANC", "ABC"}: "BANC",
	// 	{"a", "a"}:               "a",
	// 	{"a", "aa"}:              "",
	// 	{"aa", "aa"}:             "aa",
	// 	{"ab", "a"}:              "a",
	// }
	// for k, want := range m76 {
	// 	p1, p2 := k[0], k[1]
	// 	if get := interview.Q76(p1, p2); get != want {
	// 		t.Fatal("input:", p1, p2, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q88
	// i881 := [][]int{{1, 2, 3, 0, 0, 0}}
	// i882 := [][]int{{2, 5, 6}}
	// i8834 := [][2]int{{3, 3}}
	// w88 := [][]int{{1, 2, 2, 3, 5, 6}}
	// for i := 0; i < len(w88); i++ {
	// 	want, nums1, nums2, m, n := w88[i], i881[i], i882[i], i8834[i][0], i8834[i][1]
	// 	get := make([]int, len(nums1))
	// 	copy(get, nums1)
	// 	interview.Q88(nums1, m, nums2, n)
	// 	if !reflect.DeepEqual(nums1, want) {
	// 		t.Fatal("input:", get, nums2, m, n, "want:", want, "but get:", nums1)
	// 	}
	// }

	// test: q121
	// i121 := [][]int{{1, 2, 11, 4, 7}}
	// w121 := []int{10}
	// for i, prices := range i121 {
	// 	if get, want := interview.Q121(prices), w121[i]; get != want {
	// 		t.Fatal("input:", prices, "want:", want, "but get:", get)
	// 	}
	// }

	// test: q300
	// i300 := [][]int{{10, 9, 2, 5, 3, 7, 101, 18}, {0, 1, 0, 3, 2, 3}, {7, 7, 7, 7, 7, 7, 7}, {1, 3, 6, 7, 9, 4, 10, 5, 6}}
	// w300 := []int{4, 4, 1, 6}
	// for i := 0; i < len(i300); i++ {
	// 	nums, want := i300[i], w300[i]
	// 	if get := interview.Q300(nums); get != want {
	// 		t.Fatal("input:", nums, "want:", want, "but get:", get)
	// 	}
	// }

	i347 := [][]int{{2, 1, 1, 1, 2, 2, 3}, {10, 3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}}
	w347 := [][]int{{1, 2}, {1, 2, 5, 3, 6, 7, 4, 8, 10, 11}}
	for i, input := range i347 {
		k, nums, want := input[0], input[1:], w347[i]
		input := make([]int, len(nums))
		copy(input, nums)
		get := interview.Q347(nums, k)
		sort.Ints(get)
		sort.Ints(want)
		if !reflect.DeepEqual(get, want) {
			fmt.Println("input:", input, "want:", want, "but get:", get)
		}
	}
}
