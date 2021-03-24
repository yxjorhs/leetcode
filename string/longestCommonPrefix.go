package string

// LongestCommonPrefix 查找最长公共字符串
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""

	// 以str[0]作为模板与其他字符串从左到右比较每个字符是否一致
	for len(prefix) < len(strs[0]) {
		isBreak := false
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= len(prefix) {
				isBreak = true
				break
			}
			if strs[0][len(prefix)] != strs[i][len(prefix)] {
				isBreak = true
				break
			}
		}

		if isBreak {
			break
		}

		prefix += string(strs[0][len(prefix)])
	}

	return prefix
}
