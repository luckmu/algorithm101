package dpointer

import (
	"math"
)

// No.633
func judgeSquareSum(c int) bool {
	low, high := 0, int(math.Sqrt(float64(c)))
	for low <= high {
		sum := low*low + high*high
		if sum < c {
			low++
		} else if sum > c {
			high--
		} else {
			return true
		}
	}
	return false
}

// No.524
func findLongestWord(s string, dictionary []string) string {
	ret := ""
	for k := len(dictionary) - 1; k >= 0; k-- {
		i, j := 0, 0
		for i < len(s) && j < len(dictionary[k]) {
			if s[i] == dictionary[k][j] {
				j++
			}
			i++
		}
		if j >= len(dictionary[k]) {
			if len(dictionary[k]) > len(ret) {
				ret = dictionary[k]
			} else if len(dictionary[k]) == len(ret) {
				for i := 0; i < len(ret); i++ {
					if dictionary[k][i] < ret[i] {
						ret = dictionary[k]
					} else if dictionary[k][i] == ret[i] {
						continue
					} else {
						break
					}
				}
			}
		}
	}
	return ret
}
