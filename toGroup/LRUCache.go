package togroup

type lruNode struct {
	key  int
	last *lruNode
	next *lruNode
}

/*LRUCache lru缓存类*/
type LRUCache struct {
	capacity int              // 最多保存多少个key
	dict     map[int]int      // 保存key对应的值
	keyCount int              // 目前保存了多少个key
	lruHead  *lruNode         // lru链表头
	lruTail  *lruNode         // lru链表尾
	lruDict  map[int]*lruNode // 保存key对应的lru结点
}

/*Constructor 创建实例*/
func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]int{}, 0, nil, nil, map[int]*lruNode{}}
}

/*Get 获取key*/
func (t *LRUCache) Get(key int) int {
	if _, ok := t.dict[key]; ok {
		t.lruKeyMoveToHead(key)
		return t.dict[key]
	}

	return -1
}

/*Put 设置key*/
func (t *LRUCache) Put(key int, value int) {
	if _, ok := t.dict[key]; ok {
		t.keyCount--
		t.lruKeyMoveToHead(key)
	} else {
		t.lruKeyAdd(key)
	}

	t.dict[key] = value
	t.keyCount++

	if t.keyCount > t.capacity {
		delete(t.dict, t.lruTail.key)
		t.lruKeyDel(t.lruTail.key, true)
		t.keyCount--
	}
}

/*将{key}放到lru链表头部*/
func (t *LRUCache) lruKeyMoveToHead(key int) {
	if _, ok := t.lruDict[key]; !ok {
		return
	}
	t.lruKeyDel(key, false)
	t.lruKeyAdd(key)
}

/*lru链表头部添加新的结点*/
func (t *LRUCache) lruKeyAdd(key int) {
	node := t.lruDict[key]

	if node == nil {
		node = &lruNode{key, nil, nil}
		t.lruDict[key] = node
	} else {
		node.last = nil
		node.next = nil
	}

	node.next = t.lruHead

	if t.lruHead != nil {
		t.lruHead.last = node
	}

	t.lruHead = node

	if t.lruTail == nil {
		t.lruTail = node
	}
}

/*将key从lru链表中删除, {isDestoryDictKey}为true是再从字典中删除*/
func (t *LRUCache) lruKeyDel(key int, isDestoryDictKey bool) {
	node := t.lruDict[key]

	if node == nil {
		return
	}

	if node.last != nil {
		node.last.next = node.next
	}
	if node.next != nil {
		node.next.last = node.last
	}

	if t.lruHead == node {
		t.lruHead = node.next
	}

	if t.lruTail == node {
		t.lruTail = node.last
	}

	node.next = nil
	node.last = nil

	if isDestoryDictKey {
		delete(t.lruDict, key)
		// TODO 释放node占用的内存
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
