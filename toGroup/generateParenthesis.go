package togroup

type node struct {
	left int
	right int
	str string
}

/*
GenerateParenthesis 返回n个括号所有的可能组成

例:
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
*/
func GenerateParenthesis(n int) []string {
	ret := []string{}

	// 首个符号一定是"("
	nodeList := []node{
		node{ 1, 0, "(" },
	}

	// n个括号就会有n*2个字符
	for i := 0; i < n * 2 - 1; i++ {
		newNodeList := []node{}

		// 给nodeList的每个结点尝试加(或者)
		for j := 0; j < len(nodeList); j++ {
			// + (
			if nodeList[j].left < n {
				newNodeList = append(newNodeList, node{ nodeList[j].left + 1, nodeList[j].right, nodeList[j].str + "(" })
			}
			// + )
			if nodeList[j].right < nodeList[j].left {
				newNodeList = append(newNodeList, node{ nodeList[j].left, nodeList[j].right + 1, nodeList[j].str + ")" })
			}
		}

		nodeList = newNodeList
	}

	for i := 0; i < len(nodeList); i++ {
		ret = append(ret, nodeList[i].str)
	}

	return ret
}