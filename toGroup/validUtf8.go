package togroup

func validUtf8(data []int) bool {
	cou := 0

	for i := 0; i < len(data); i++ {
		if cou == 0 {
			len := headCheck(data[i])
			if len == 0 {
				return false
			}
			cou = len - 1
		} else {
			if bodyCheck(data[i]) == false {
				return false
			}
			cou--
		}
	}

	if cou != 0 {
		return false
	}

	return true
}

/*检查是否是头字符串，返回0校验失败，返回值为字符长度*/
func headCheck(v int) int {
	len := 0

	// 0xxxxxxx
	if v>>7 == 0 {
		return 1
	}

	for i := 0; i < 3; i++ {
		mask := 255 >> (4 + i) << 1 // 110, 1110, 11110

		if v>>(3+i) == mask {
			len = 4 - i
			break
		}
	}

	return len
}

func bodyCheck(v int) bool {
	return v>>6 == 2
}
