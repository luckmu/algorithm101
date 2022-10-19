package test

import (
	"algorithm101/interview"
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

	m7 := map[int]int{
		123:  321,
		-123: -321,
		120:  21,
	}
	for input, want := range m7 {
		get := interview.Q7(input)
		if want != get {
			t.Fatal("input:", input, "want:", want, "but get:", get)
		}
	}

}
