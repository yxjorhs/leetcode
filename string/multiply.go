package string

const digits = "0123456789"

// Multiply 字符串相乘
func Multiply(num1 string, num2 string) string {
	if len(num1) == 1 && len(num2) == 1 {
		return intToStr(byteToInt(num1[0]) * byteToInt(num2[0]))
	}

	ret := "0"

	for i, zero := 0, ""; i < len(num1); i++ {
		for j, zero2 := 0, ""; j < len(num2); j++ {
			ret = Add(ret, Multiply(string(num1[len(num1)-1-i]), string(num2[len(num2)-1-j]))+zero+zero2)
			zero2 += "0"
		}

		zero += "0"
	}

	return rmUselessZero(ret)
}

func byteToInt(b byte) int {
	return int(b - 48)
}

func intToStr(i int) string {
	ret := string(digits[i%10])
	i = i / 10

	for i > 0 {
		ret = string(digits[i%10]) + ret
		i = i / 10
	}

	return ret
}

func rmUselessZero(str string) string {
	ret := ""

	for i := 0; i < len(str); i++ {
		if ret == "" && string(str[i]) == "0" && i != len(str)-1 {
			continue
		}

		ret += string(str[i])
	}

	return ret
}
