package interview

// time 46.90%; mem 12.99%
func minWindow(s string, t string) string {
	// minimum, 最小 O(n)
	if len(s) < len(t) {
		return ""
	}
	need := make(map[byte]int)
	have := make(map[byte]int)
	needchs := make(map[byte]struct{})
	for i := 0; i < len(t); i++ {
		need[t[i]]++
		needchs[t[i]] = struct{}{}
	}
	var l, r int
	var ret string = ""
	var needmore bool = true
	for l <= r && r < len(s) {
		for ; needmore && r < len(s); r++ {
			// r++ until needmore=f
			ch := s[r]
			if _, ok := need[ch]; ok {
				have[ch]++
				if have[ch] >= need[ch] {
					delete(needchs, ch)
					if len(needchs) == 0 {
						needmore = false
					}
				}
			}
		}
		// loop -> needmore=false then r++, so last index is r-1
		for ; !needmore && l <= r-1; l++ {
			// l++ until needmore=t
			// cover outside r, `r := r-1`
			ch, r := s[l], r-1
			if _, ok := need[ch]; ok {
				have[ch]--
				if have[ch] < need[ch] {
					needchs[ch] = struct{}{}
					needmore = true
				}
			}
			if ret == "" || r-l+1 < len(ret) {
				ret = s[l : r+1]
			}
		}
	}
	return ret
}

func Q76(s, t string) string {
	return minWindow(s, t)
}
