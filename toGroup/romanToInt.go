package togroup

// RomanToInt 罗马转数字
func RomanToInt(s string) int {
	romanIntMap := map[string]int {
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	ret := 0

	for i:=0; i< len(s); i++ {
		if (i > 0 && romanIntMap[string(s[i])] > romanIntMap[string(s[i - 1])]) {
			ret += romanIntMap[string(s[i])]
			ret -= romanIntMap[string(s[i - 1])] * 2
		} else {
			ret += romanIntMap[string(s[i])]
		}
	}

	return ret
}
