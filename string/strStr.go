package string

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		if haystack == needle {
			return 0
		}

		return -1
	}

	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		isMatch := true

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				isMatch = false
				break
			}
		}

		if isMatch == true {
			return i
		}
	}

	return -1
}
