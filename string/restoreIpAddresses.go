package string

func restoreIpAddresses(s string) []string {
	if len(s) < 4 && len(s) > 12 {
		return []string{}
	}

	ret := []string{}

	for i := 1; i <= len(s)-1 && i <= 3; i++ {
		if isValidNum(string(s[:i])) == false {
			continue
		}
		for j := i + 1; j <= len(s)-1 && j <= i+3; j++ {
			if isValidNum(string(s[i:j])) == false {
				continue
			}
			for k := j + 1; k <= len(s)-1 && k <= j+3; k++ {
				if isValidNum(string(s[j:k])) == true && isValidNum(string(s[k:])) == true {
					ret = append(ret, string(s[:i])+"."+string(s[i:j])+"."+string(s[j:k])+"."+string(s[k:]))
				}
			}
		}
	}

	return ret
}

func isValidNum(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}

	if string(s[0]) == "0" && len(s) != 1 {
		return false
	}

	if len(s) <= 2 {
		return true
	}

	return str2Num(s) <= 255
}

func str2Num(s string) int {
	if len(s) == 1 {
		return int(s[0] - 48)
	}

	ret := 0

	for i := 0; i < len(s); i++ {
		ret = str2Num(string(s[i])) + ret*10
	}

	return ret
}
