package interview

func rslide(s string, r, prev int) string {
	str := ""
	for l := r - prev; l >= 0; l-- {
		f, i, j := true, l, r
		for i <= j {
			if s[i] != s[j] {
				f = false
				break
			}
			i++
			j--
		}
		// s[l:r] is palindrome
		if f && len(s[l:r+1]) > len(str) {
			str = string(s[l : r+1])
		}
	}
	return str
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	subs := ""
	for r := 0; r < len(s); r++ {
		if str := rslide(s, r, len(subs)); len(str) > len(subs) {
			subs = str
		}
	}
	return subs
}

func Q5(s string) string {
	// runtime less than 38.33%
	// memory less than 89.38%
	return longestPalindrome(s)
}
