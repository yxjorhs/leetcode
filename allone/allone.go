package allone

/*
请你实现一个数据结构支持以下操作：

Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否则使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任何事情。key 保证不为空字符串。
GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串"" 。
GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
*/

type linkNode struct {
	key  string
	last *linkNode
	next *linkNode
}

type dictEle struct {
	val *int
	ptr *linkNode
}

/*AllOne AllOne*/
type AllOne struct {
	dict     map[string]dictEle
	linkHead *linkNode
	linkTail *linkNode
}

/*Constructor Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{map[string]dictEle{}, nil, nil}
}

/*Inc Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (t *AllOne) Inc(key string) {
	if e, ok := t.dict[key]; ok {
		(*e.val)++
		// 结点前移
		for t.linkNodeMoveFrontIfNeed(e.ptr) == 1 {
		}
	} else {
		v := 1
		n := &linkNode{key, nil, nil}
		t.dict[key] = dictEle{&v, n}
		t.linkNodeInsert(n, t.linkTail, nil)
		t.linkTail = n
		if t.linkHead == nil {
			t.linkHead = n
		}
	}
}

/*Dec Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (t *AllOne) Dec(key string) {
	if e, ok := t.dict[key]; ok {
		(*e.val)--

		if (*e.val) == 0 {
			elast := e.ptr.last
			enext := e.ptr.next
			t.linkNodeDel(e.ptr)
			if t.linkHead == e.ptr {
				t.linkHead = enext
			}

			if t.linkTail == e.ptr {
				t.linkTail = elast
			}
			delete(t.dict, key)
		} else {
			// 结点后移
			for t.linkNodeMoveFrontIfNeed(e.ptr.next) == 1 {
			}
		}
	}
}

/*GetMaxKey Returns one of the keys with maximal value. */
func (t *AllOne) GetMaxKey() string {
	if t.linkHead == nil {
		return ""
	}

	return t.linkHead.key
}

/*GetMinKey Returns one of the keys with Minimal value. */
func (t *AllOne) GetMinKey() string {
	if t.linkTail == nil {
		return ""
	}

	return t.linkTail.key
}

/*假设结点的值小于前面结点的值，则将结点与前面的结点交换位置，返回0表示交换失败，1则交换成功*/
func (t *AllOne) linkNodeMoveFrontIfNeed(n *linkNode) int {
	if n == nil || n.last == nil {
		return 0
	}

	if (*t.dict[n.key].val) <= (*t.dict[n.last.key].val) {
		return 0
	}

	left := n.last.last
	right := n.last
	t.linkNodeDel(n)
	t.linkNodeInsert(n, left, right)

	if t.linkTail == n {
		t.linkTail = n.next
	}

	if t.linkHead == n.next {
		t.linkHead = n
	}

	return 1
}

/*删除{n},链接{n}的上下结点*/
func (t *AllOne) linkNodeDel(n *linkNode) {
	if n == nil {
		return
	}

	if n.last != nil {
		n.last.next = n.next
	}

	if n.next != nil {
		n.next.last = n.last
	}

	n.next = nil
	n.last = nil
}

/*在{last}与{next}中间插入新结点{n}*/
func (t *AllOne) linkNodeInsert(n *linkNode, last *linkNode, next *linkNode) {
	if last != nil {
		last.next = n
		n.last = last
	}

	if next != nil {
		n.next = next
		next.last = n
	}
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
