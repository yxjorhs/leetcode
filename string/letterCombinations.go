package string

// import "fmt"

/*
LetterCombinations 给定2-9的数字，返回他们在9键上所代表的字母所能组成的字符

例:
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/
func LetterCombinations(digits string) []string {
	ret := []string{}

	// 字典 - 数字对应的英文字符
	numCharMap := map[string][]string {
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	for i := 0; i < len(digits); i++ {
		// 保存此时的ret长度
		lenStore := len(ret);

		if lenStore == 0 {
			// ret长度为0时将数字对应的字符丢到ret中
			for k := 0; k < len(numCharMap[string(digits[i])]); k++ {
				ret = append(ret, numCharMap[string(digits[i])][k])
			}
			continue
		}

		// ret长度不为0时，给原来所有的元素添加后缀
		for j:= 0; j < lenStore; j++ {
			char := ret[j]
			for k := 0; k < len(numCharMap[string(digits[i])]); k++ {
				if k == 0 {
					// 旧元素加后缀
					ret[j] = char + numCharMap[string(digits[i])][k]
					continue
				}

				// 多个后缀则再增加新的元素
				ret = append(ret, char + numCharMap[string(digits[i])][k])
			}
		}
	}

	return ret
}