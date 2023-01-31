package dfs

var _ = letterCombinationsDFS

var keymap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func dfs17(combine []byte, digits string, ret *[]string) {
	if len(digits) == 0 {
		if len(combine) != 0 {
			*ret = append(*ret, string(combine))
		}
		return
	}
	chs := keymap[digits[0]]
	for i := 0; i < len(chs); i++ {
		dfs17(append(combine, chs[i]), digits[1:], ret)
	}
}

func letterCombinationsDFS(digits string) []string {
	ret := []string{}
	dfs17(make([]byte, 0), digits, &ret)
	return ret
}
