package string

// CheckInclusion 字符串的排列，第一个字符串的排列之一是第二个字符串的子串
// 例: s1 = "ab", s2 = "eiodbaooo", return true
func CheckInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	s1dict := buildDict(s1) // s1构建字典

	ret := false

	s2dict := map[string]int{} // s2字串字典

	for left, right := 0, 0; right < len(s2); right++ {
		if s1dict[string(s2[right])] == 0 { // s2出现s1不存在的字符
			left = right + 1
			s2dict = map[string]int{}
			continue
		}

		s2dict[string(s2[right])]++

		if right-left+1 < len(s1) { // s2子串长度与s1长度不同
			continue
		}

		if compareDict(s1dict, s2dict) { // s1字典与s2子串字典相同则说明s1的排列只一是s2的字串
			ret = true
			break
		}

		s2dict[string(s2[left])]--
		left++
	}

	return ret
}

func buildDict(s string) map[string]int {
	dict := map[string]int{}

	for i := 0; i < len(s); i++ {
		dict[string(s[i])]++
	}

	return dict
}

func compareDict(d1 map[string]int, d2 map[string]int) bool {
	for k, v := range d1 {
		if d2[k] != v {
			return false
		}
	}

	return true
}
