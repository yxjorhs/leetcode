package string

type stack struct {
	Data []string
}

func (this *stack) in(str string) {
	this.Data = append(this.Data, str)
}

func (this *stack) out() *string {
	if len(this.Data) == 0 {
		return nil
	}

	v := this.Data[len(this.Data)-1]
	this.Data = this.Data[:len(this.Data)-1]
	return &v
}

func reverseWords(s string) string {
	stk := new(stack)
	word := ""

	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			word += string(s[i])
		}

		if string(s[i]) == " " || i == len(s)-1 {
			if len(word) > 0 {
				stk.in(word)
			}
			word = ""
		}
	}

	ret := ""

	for {
		v := stk.out()

		if v == nil {
			break
		}

		if len(ret) != 0 {
			ret += " "
		}

		ret += *v
	}

	return ret
}
