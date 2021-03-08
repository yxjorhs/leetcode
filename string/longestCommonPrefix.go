package string

// LongestCommonPrefix 查找最长公共字符串
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	str1 := strs[0]

	str1MatchLen := len(str1)

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < str1MatchLen {
			str1MatchLen = len(strs[i])
		}

		for j := 0; j < str1MatchLen && j < len(strs[i]); j++ {
			if strs[i][j] != str1[j] {
				str1MatchLen = j
			}
		}
	}

	prefix := ""

	for i := 0; i < str1MatchLen; i++ {
		prefix += string(str1[i])
	}

	return prefix
}