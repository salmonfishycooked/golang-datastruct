package datastruct

import "fmt"

type linkedListNode struct {
	data int             // 数据域
	next *linkedListNode // 存储后继结点存储位置，若为空则该结点为表尾结点
}

// 单链表结构体
type LinkedList struct {
	head   *linkedListNode // 存放头结点
	length int             // 存储链表长度
}

func NewLinkedList() *LinkedList {
	node := &linkedListNode{ // 头结点
		data: -1,
		next: nil,
	}

	return &LinkedList{
		head:   node,
		length: 0,
	}
}

// insert 在第 i 个位置之前插入新的结点
func (l *LinkedList) insert(position int, data int) {
	if position > l.length || position < 0 {
		return
	}

	node := &linkedListNode{
		data: data,
		next: nil,
	}
	ptr := l.head
	for i := 0; i < position; i++ { // 寻找需插入结点的前一个结点存储位置
		ptr = l.head.next
	}

	node.next = ptr.next
	ptr.next = node

	l.length++
}

// print 打印链表中所有数据
func (l *LinkedList) print() {
	ptr := l.head.next
	for i := 0; i < l.length; i++ {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}
