package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 0
	fmt.Scan(&input)

	for i := 2; i <= 16; i++ {
		v := numToR(input, i)
		if v == -1 {
			continue
		}
		if isHW(v) {
			fmt.Println(i)
		}
	}
}

func numToR(n int, r int) int {
	if r < 2 || r > 16 {
		return -1
	}

	ret := 0
	temp := 1

	for n > 0 {
		mod := n % r
		n = n / r
		ret += temp * mod
		temp *= 10
	}

	return ret
}

// isHW 是否回文
func isHW(n int) bool {
	str := strconv.Itoa(n)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func search(arr []int, diff int) [][]int {
	ret := [][]int{}
	// 字典， arr[?] - diff => 下标
	dict := map[int][]int{}
	for i := 0; i < len(arr); i++ {
		list := dictSet(&dict, diff, arr[i], i)
		if len(list) > 0 {
			for j := 0; j < len(list); j++ {
				ret = append(ret, []int{i, list[j]})
			}
		}
	}

	dict = map[int][]int{}
	for i := len(arr) - 1; i >= 0; i-- {
		list := dictSet(&dict, diff, arr[i], i)
		if len(list) > 0 {
			for j := 0; j < len(list); j++ {
				ret = append(ret, []int{i, list[j]})
			}
		}
	}

	return ret
}

func dictSet(dict *map[int][]int, diff int, v int, idx int) []int {
	ret := []int{}
	if _, ok := (*dict)[v]; ok {
		ret = (*dict)[v]
	}

	if _, ok := (*dict)[v-diff]; ok {
		// println("e", v, idx)
		(*dict)[v-diff] = append((*dict)[v-diff], idx)
	} else {
		// println("f", v, idx)

		(*dict)[v-diff] = []int{idx}
	}

	return ret
}

// // 本题为考试单行多行输入输出规范示例，无需提交，不计分。
// package main

// import (
//     "fmt"
// )
// func main() {
//     a:=0
//     b:=0
//     for {
//         n, _ := fmt.Scan(&a,&b)
//         if n == 0 {
//             break
//         } else {
//             fmt.Printf("%d\n",a+b)
//         }
//     }
// }

// 本题为考试多行输入输出规范示例，无需提交，不计分。
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	n := 0
// 	ans := 0

// 	fmt.Scan(&n)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			x := 0
// 			fmt.Scan(&x)
// 			ans = ans + x
// 		}
// 	}
// 	fmt.Printf("%d\n", ans)
// }
