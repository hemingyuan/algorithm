package base

type LinkListNode struct {
	value int
	next  *LinkListNode
}

func (node *LinkListNode) IsEqual(target *LinkListNode) bool {
	return node.value == target.value && node.next == target.next
}

type LinkList struct {
	head   *LinkListNode
	length int
}

func CreateLinkList() *LinkList {
	return &LinkList{head: nil, length: 0}
}

// Insert 插入到头部
func (ll *LinkList) Insert(value int) *LinkListNode {
	node := LinkListNode{value: value}
	if ll.head != nil {
		node.next = ll.head
	}
	ll.head = &node
	ll.length++
	return &node
}

func (ll *LinkList) InsertAfter(value int, pos *LinkListNode) {
	node := LinkListNode{value: value}
	if pos.next != nil {
		node.next = pos.next
	}
	pos.next = &node
	ll.length++
}

func (ll *LinkList) Find(value int) *LinkListNode {
	for cur := ll.head; cur != nil; cur = cur.next {
		if value == cur.value {
			return cur
		}
	}
	return nil
}

// Delete 从头部删除
func (ll *LinkList) Delete() *LinkListNode {
	if ll.head == nil {
		return nil
	}
	node := ll.head
	ll.head = node.next
	ll.length--
	return node
}

// 根据取值删除
func (ll *LinkList) DeleteByValue(value int) *LinkListNode {
	cur := ll.head
	pre := ll.head
	for ; cur != nil; cur = cur.next {
		if value == cur.value {
			ll.length--
			if pre == ll.head {
				ll.head = nil
				return pre
			} else {
				pre.next = cur.next
				return cur
			}
		}
		pre = cur
	}
	return nil
}

// Visit 遍历高阶函数
func (ll *LinkList) Visit(fn func(node *LinkListNode)) {
	for cur := ll.head; cur != nil; cur = cur.next {
		fn(cur)
	}
}

// HasLoop 是否存在闭环，两种实现
// 1 增加一个外部存储标记已遍历的节点看是否有重复
// 2 使用两个指针，采用不同的步进长度，如果两者相遇则表示有环路
func (ll *LinkList) HasLoop() bool {
	c1 := ll.head
	if c1 == nil {
		return false
	}
	c2 := c1.next
	for ; c1 != nil; c1 = c1.next {
		if c1.IsEqual(c2) {
			return true
		}
		if c2 == nil {
			return false
		}
		if c2.next == nil {
			return false
		}
		c2 = c2.next.next
	}
	return false
}

// MergeSortedLinkList 合并有序链表
// 1: 归并排序，开一个新的链表，每次将小的放进去
// 2: 原地归并，使用其中一条作为基线进行归并，找到两者的小值之后，再在基线上移动游标寻找插入位置
func MergeSortedLinkList(l1 *LinkList, l2 *LinkList) *LinkList {
	c1 := l1.head
	c2 := l2.head

	// // 普通归并
	// head := &LinkListNode{}
	// c := head
	// for c1 != nil && c2 != nil {
	// 	if c1.value < c2.value {
	// 		c.next = &LinkListNode{value: c1.value}
	// 		c = c.next
	// 		c1 = c1.next
	// 	} else {
	// 		c.next = &LinkListNode{value: c2.value}
	// 		c = c.next
	// 		c2 = c2.next
	// 	}
	// }
	// if c1 != nil {
	// 	c.next = c1
	// }
	// if c2 != nil {
	// 	c.next = c2
	// }
	// return &LinkList{length: l1.length + l2.length, head: head.next}

	// 空间优化解法，基准是l1
	pos := l1.head
	for c1 != nil && c2 != nil {
		if c1.value < c2.value {
			if pos == c1 {
				continue
			}
			pos.next = c1
			pos = pos.next
			c1 = c1.next
			continue
		}
		if pos == l1.head {
			l1.head = c2
			pos = l1.head
			c2 = c2.next
			continue
		}
		pos.next = c2
		pos = pos.next
		c2 = c2.next
	}
	if c1 != nil {
		pos.next = c1
	}
	if c2 != nil {
		pos.next = c2
	}

	return &LinkList{length: l1.length + l2.length, head: l1.head}
}

type DualLinkList struct {
	head   *DualLinkListNode
	tail   *DualLinkListNode
	length int
}

type DualLinkListNode struct {
	pre   *DualLinkListNode
	next  *DualLinkListNode
	value int
}

func CreateDualLinkList() *DualLinkList {
	return &DualLinkList{head: nil, tail: nil, length: 0}
}

// Insert 插入到头部
func (dll *DualLinkList) Insert(value int) {
	node := DualLinkListNode{value: value}
	if dll.length == 0 {
		dll.head = &node
		dll.tail = &node
		dll.length++
		return
	}
	// 头部处理
	dll.head.pre = &node
	node.next = dll.head
	dll.head = &node
	dll.length++
}

// Append 追加到尾部
func (dll *DualLinkList) Append(value int) {
	node := DualLinkListNode{value: value}
	if dll.length == 0 {
		dll.tail = &node
		dll.head = &node
		dll.length++
		return
	}
	dll.tail.next = &node
	node.pre = dll.tail
	dll.tail = &node
	dll.length++
}

func (dll *DualLinkList) Find(value int) *DualLinkListNode {
	for cur := dll.head; cur != nil; cur = cur.next {
		if value == cur.value {
			return cur
		}
	}
	return nil
}

// Pop 从头部删除
func (dll *DualLinkList) Pop() *DualLinkListNode {
	if dll.length == 0 {
		return nil
	}
	if dll.length == 1 {
		node := dll.head
		dll.head = nil
		dll.tail = nil
		return node
	}
	node := dll.head
	dll.head = node.next
	dll.head.pre = nil
	dll.length--
	return node
}

// Shift 从尾部删除
func (dll *DualLinkList) Shift() *DualLinkListNode {
	if dll.length == 0 {
		return nil
	}
	if dll.length == 1 {
		node := dll.head
		dll.head = nil
		dll.tail = nil
		return node
	}
	node := dll.tail
	dll.tail = node.pre
	dll.tail.next = nil
	dll.length--
	return node
}

// DeleteByValue 根据取值删除
func (dll *DualLinkList) DeleteByValue(value int) *DualLinkListNode {
	for cur := dll.head; cur != nil; cur = cur.next {
		if value == cur.value {
			if cur.pre == nil {
				dll.head = dll.head.next
				if dll.head != nil {
					dll.head.pre = nil
				}
			} else {
				cur.pre.next = cur.next
			}
			if cur.next == nil {
				dll.tail = dll.tail.pre
				if dll.tail != nil {
					dll.tail.next = nil
				}
			} else {
				cur.next.pre = cur.pre
			}
			dll.length--
			return cur
		}
	}
	return nil
}

// Visit 遍历高阶函数
func (dll *DualLinkList) Visit(fn func(node *DualLinkListNode)) {
	for cur := dll.head; cur != nil; cur = cur.next {
		fn(cur)
	}
}
