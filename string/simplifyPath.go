package string

func simplifyPath(path string) string {
	stk := new(stack)

	temp := ""

	for i := 0; i < len(path); i++ {
		if string(path[i]) != "/" {
			temp += string(path[i])
		}

		if string(path[i]) == "/" || i == len(path)-1 {
			if temp == ".." {
				stk.out()
			} else if temp == "." || temp == "" {

			} else {
				stk.in(temp)
			}

			temp = ""
		}
	}

	ret := ""

	for {
		v := stk.out()

		if v == nil {
			break
		}

		ret = "/" + *v + ret
	}

	if ret == "" {
		return "/"
	}

	return ret
}
