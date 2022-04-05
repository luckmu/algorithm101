package dpointer

// 680
// runtime 78.35%; memory 99.66%
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			if s[l] == s[r-1] {
				if r := absPalindrom(s[l:r]); r {
					return r
				}
			}
			if s[l+1] == s[r] {
				if r := absPalindrom(s[l+1 : r+1]); r {
					return r
				}
			}
			return false
		}
		l, r = l+1, r-1
	}
	return true
}

func absPalindrom(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
