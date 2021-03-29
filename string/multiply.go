package string

import (
	"strconv"
)

// Multiply 字符串相乘
func Multiply(num1 string, num2 string) string {
	// 当两个字符串的长度小于calLen时直接转为整数计算
	calLen := 8

	if len(num1) <= calLen && len(num2) <= calLen {
		v1, _ := strconv.Atoi(num1)
		v2, _ := strconv.Atoi(num2)
		return strconv.Itoa(v1 * v2)
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
