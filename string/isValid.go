package string

/*
IsValid 校验s中的括号是否有效
*/
func IsValid(s string) bool {
	leftMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	stack := []string{}

	for i := 0; i < len(s); i++ {
		char := string(s[i])

		if _, ok := leftMap[char]; ok {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if leftMap[stack[len(stack)-1]] == char {
			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
