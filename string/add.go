package string

import (
	"strconv"
)

// Add 字符串相加
func Add(num1 string, num2 string) string {
	longer := num1
	shorter := num2

	if len(num2) > len(num1) {
		longer = num2
		shorter = num1
	}

	longer = rmUselessZero(longer)
	shorter = rmUselessZero(shorter)

	incr := 0

	for i := 0; i < len(longer); i++ {
		sum := int(longer[len(longer)-1-i] - 48)

		sum += incr

		if len(shorter)-1-i >= 0 {
			sum += int(shorter[len(shorter)-1-i] - 48)
		}

		longer = longer[:len(longer)-1-i] + strconv.Itoa(sum%10) + longer[len(longer)-i:]

		incr = sum / 10
	}

	if incr > 0 {
		longer = strconv.Itoa(incr) + longer
	}

	return longer
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
