package togroup

// IntToRoman 数字转罗马数字
func IntToRoman(num int) string {
	intRomanMap := map[int]string {
		1: "I",
		5: "V",
		10: "X",
		50: "L",
		100: "C",
		500: "D",
		1000: "M",
	}
	ret := ""

	arr := []int{1000, 100, 10, 1}

	for i:=0; i < len(arr); i++ {
		mod := arr[i]
		div := num / mod
		num = num % mod

		if (div == 0) {
			continue
		} else if (div <= 3) {
			for j:=0; j < div; j++ {
				ret += intRomanMap[mod]
			}
		} else if (div == 4) {
			ret += intRomanMap[mod]
			ret += intRomanMap[mod * 5]
		} else if (div == 5) {
			ret += intRomanMap[mod * 5]
		} else if (div <= 8) {
			ret += intRomanMap[mod * 5]
			for j:=0; j < div - 5; j++ {
				ret += intRomanMap[mod]
			}
		} else if (div == 9) {
			ret += intRomanMap[mod]
			ret += intRomanMap[mod * 10]
		}
	}

	return ret
}
