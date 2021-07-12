package stack

/*MinStack 最小栈*/
type MinStack struct {
	list []int
	min  int
}

/*Constructor initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

/*Push 入栈*/
func (t *MinStack) Push(val int) {
	t.list = append(t.list, val)

	if (val < t.min) || (len(t.list) == 1) {
		t.min = val
	}
}

/*Pop 出栈*/
func (t *MinStack) Pop() {
	if len(t.list) == 0 {
		return
	}

	if len(t.list) == 1 {
		t.list = []int{}
		t.min = 0
		return
	}

	t.list = t.list[:len(t.list)-1]

	t.min = t.list[0]

	for i := 0; i < len(t.list); i++ {
		if t.list[i] < t.min {
			t.min = t.list[i]
		}
	}
}

/*Top 返回栈顶指*/
func (t *MinStack) Top() int {
	return t.list[len(t.list)-1]
}

/*GetMin 返回栈最小值*/
func (t *MinStack) GetMin() int {
	return t.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
