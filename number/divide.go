package number

const maxInt = 2147483647

func divide(dividend int, divisor int) int {
	arr := []int{} // 二进制数组，arr[0]代表2的0次幂是1

	for 

	if abs(dividend) < abs(divisor) {
		return 0
	}

	sum := divisor
	cou := 1

	if dividend>>31 != divisor>>31 {
		cou = -1
	}

	for {
		if abs(sum+sum) > abs(dividend) {
			break
		}

		sum += sum
		cou += cou
	}

	leave := abs(dividend) - abs(sum)

	if dividend < 0 {
		leave = 0 - leave
	}

	ret := cou + divide(leave, divisor)

	if ret > maxInt {
		return maxInt
	}

	return ret
}

func abs(num int) int {
	if num < 0 {
		return 0 - num
	}

	return num
}
