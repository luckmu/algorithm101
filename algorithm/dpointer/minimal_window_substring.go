package dpointer

// runtime 53.89%, memory 84.9%
func minWindow(s string, t string) string {
	ret, origlen := "", 0
	// why origlen is length of t or s ?
	// s is easy to understand but t ?
	// 
	if len(s) > len(t) {
		origlen = len(s)
	} else {
		origlen = len(t)
	}
	need, needchs := make(map[byte]int), make(map[byte]struct{})
	for i := 0; i < len(t); i++ {
		need[t[i]]++
		needchs[t[i]] = struct{}{}
	}
	l, r, own := 0, 0, make(map[byte]int)
	for r < len(s) || len(needchs) == 0 {
		// extend right until len(needchs) == 0
		for r < len(s) && len(needchs) != 0 {
			if _, ok := need[s[r]]; ok {
				own[s[r]]++
				if own[s[r]] >= need[s[r]] {
					delete(needchs, s[r])
				}
			}
			r++
		}
		if origlen >= len(s[l:r]) && len(needchs) == 0 {
			ret = s[l:r]
			origlen = len(ret)
		}

		// extend left until len(needchs) != 0
		for l < r && len(needchs) == 0 {
			if _, ok := need[s[l]]; ok {
				own[s[l]]--
				if own[s[l]] < need[s[l]] {
					needchs[s[l]] = struct{}{}
				}
			}
			if len(needchs) == 0 && origlen >= len(s[l+1:r]) {
				ret = s[l+1 : r]
				origlen = len(ret)
			}
			l++
		}
	}
	return ret
}
