package string

// LengthOfLongestSubstring 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	maxLen := 0

	// 使用两个指针标记当前检查的字符串的头尾部，right不断读取新的字符，left检查left到right之间有没有重复的字符
	for left, right := 0, 0; right < len(s); right++ {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				left = i + 1
				break
			}
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
