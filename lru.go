package main

type lruNode struct {
	key  int
	next *lruNode
	last *lruNode
}

type dictNode struct {
	val    int
	lruPtr *lruNode
}

type LRUCache struct {
	size    int
	use     int
	dict    map[int]*dictNode
	lruHead *lruNode
	lruTail *lruNode
}

func LRUCacheNew(size int) *LRUCache {
	return &LRUCache{
		size,
		0,
		map[int]*dictNode{},
		nil,
		nil,
	}
}

func (t *LRUCache) Put(key int, val int) {
	// println("input", key)
	lruPtr := &lruNode{}
	if v, ok := t.dict[key]; ok {
		v.val = val
		lruPtr = v.lruPtr
	} else {
		// 容量不足，删除lru链表最后一个
		if t.use >= t.size {
			// println("del", t.lruTail.key)
			last := t.lruTail.last
			lruMove(t.lruTail, nil, nil)
			delete(t.dict, t.lruTail.key)
			t.lruTail = last
			t.use--
		}
		lruPtr.key = key
		node := dictNode{
			val,
			lruPtr,
		}
		t.dict[key] = &node
		t.use++
	}

	t.lruHeadSet(lruPtr)
}

func (t *LRUCache) Get(key int) int {
	if v, ok := t.dict[key]; ok {
		t.lruHeadSet(v.lruPtr)
		return v.val
	} else {
		return -1
	}
}

// 移动节点{ptr}到lru链表头
func (t *LRUCache) lruHeadSet(ptr *lruNode) {
	// 加入ptr是lru链表尾部需要更新lruTail
	if t.lruTail == ptr {
		t.lruTail = ptr.last
	}

	lruMove(ptr, nil, t.lruHead)
	t.lruHead = ptr
	if t.lruTail == nil {
		t.lruTail = ptr
	}
}

func lruMove(target *lruNode, last *lruNode, next *lruNode) {
	if last != nil {
		last.next = target
	}
	if next != nil {
		next.last = target
	}
	target.last = last
	target.next = next
}
