package array

/*
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。



示例 1：

输入：n = 3, k = 3
输出："213"
示例 2：

输入：n = 4, k = 9
输出："2314"
示例 3：

输入：n = 3, k = 1
输出："123"


提示：

1 <= n <= 9
1 <= k <= n!

解法:
输出的长度即为n
再根据k推算第1,2,3...n位的值，例

当n=3, k=4时:

store := [1,2,3]

第1位:
	mod = (n - 1)!
	idx = (k - 1) / mod
	val = store[idx]
	k -= idx * mod
	将val从store中移出
第2位:
	mod = (n - 2)!
	idx = (k - 1) / mod
	val = store[idx]
	k -= idx * mod
	将val从store中移出
第3位:
	此时store只剩下一个元素
	v3 = store[0]
*/
func getPermutation(n int, k int) string {
	store := make([]int, n)
	factorialArr := make([]int, n)
	ret := ""

	for i := 0; i < n; i++ {
		store[i] = i + 1
	}

	factorialGet(n, &factorialArr)

	for i := 0; i < n; i++ {
		if i == n-1 {
			ret += numToStr(store[0])
			break
		}

		mod := factorialArr[n-i-2]

		idx := (k - 1) / mod

		ret += numToStr(store[idx])

		k -= idx * mod

		store = append(store[:idx], store[idx+1:]...)
	}

	return ret
}

/* factorialGet 返回1到{n}的阶乘数组，[1!,2!,...n!] */
func factorialGet(n int, arr *[]int) {
	tmp := 1

	for i := 0; i < n; i++ {
		tmp *= i + 1
		(*arr)[i] = tmp
	}

}

func numToStr(n int) string {
	return string(rune(n) + 48)
}
