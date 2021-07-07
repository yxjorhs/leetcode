package number

const maxInt = 2147483647

func divide(dividend int, divisor int) int {
	resSign := 1

	if dividend>>31 != divisor>>31 {
		resSign = -1
	}

	res := 0
	dividend = abs(dividend)
	divisor = abs(divisor)

	arr := []int{} // 二进制数组，arr[0] = 10代表(2^0)个divisor的和是10
	for i, sum := 0, divisor; sum <= dividend; i += i {
		arr = append(arr, sum)
		sum += sum
	}

	for dividend >= divisor {
		divisorCount := 1
		sum := arr[0]

		for i := 1; i < len(arr); i++ {
			if arr[i] > dividend {
				break
			}
			divisorCount += divisorCount
			sum = arr[i]
		}

		res += divisorCount
		dividend -= sum
	}

	if resSign == -1 {
		return 0 - res
	}

	if res > maxInt {
		return maxInt
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return 0 - num
	}

	return num
}
