package string

// Multiply 字符串相乘
func Multiply(num1 string, num2 string) string {
	// 当两个字符串的长度小于calLen时直接转为整数计算
	calLen := 8

	if len(num1) <= calLen && len(num2) <= calLen {
		return i2s(s2i(num1) * s2i(num2))
	}

	ret := "0"

	for n1right, zero := len(num1)-1, ""; n1right >= 0; n1right -= calLen {
		v1 := num1[:n1right+1]

		n1left := n1right + 1 - calLen

		if n1left > 0 {
			v1 = v1[n1left:]
		}

		for n2right, zero2 := len(num2)-1, ""; n2right >= 0; n2right -= calLen {
			v2 := num2[:n2right+1]

			n2left := n2right + 1 - calLen

			if n2left > 0 {
				v2 = v2[n2left:]
			}

			ret = Add(ret, Multiply(v1, v2)+zero+zero2)

			zero2 += "00000000"
		}

		zero += "00000000"
	}

	return ret
}

func s2i(s string) int {
	dict := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	ret := 0

	for i := 0; i < len(s); i++ {
		ret += dict[string(s[i])]

		if i != len(s)-1 {
			ret *= 10
		}
	}

	return ret
}

func i2s(i int) string {
	digits := "0123456789"
	ret := string(digits[i%10])
	i = i / 10

	for i > 0 {
		ret = string(digits[i%10]) + ret
		i = i / 10
	}

	return ret
}
