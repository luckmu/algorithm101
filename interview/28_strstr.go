package interview

// runtime 100.00%; mem 69.83%
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		equal := true
		for j := 0; j < len(needle); j++ {
			if i+j > len(haystack)-1 ||
				haystack[i+j] != needle[j] {
				equal = false
				break
			}
		}
		if equal {
			return i
		}
	}
	return -1
}

func Q28(haystack, needle string) int {
	return strStr(haystack, needle)
}
