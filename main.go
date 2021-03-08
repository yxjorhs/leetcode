package main

import (
	"fmt"

	"github.com/yxjorhs/leetcode/tree"
)

func main() {
	midorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	t := tree.BuildByInAndPost(midorder, postorder)
	fmt.Println(tree.IsBST(t))
}
