package leet

// var strs []string = []string{"flower", "flow", "flight"}
// var strs []string = []string{"dog", "racecar", "car"}
// var strs []string = []string{"cir", "car"}
func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for _, v := range strs {
			if i >= len(v) || v[i] != ch {
				return prefix
			}
		}
		prefix += string(ch)
	}
	return prefix
}
