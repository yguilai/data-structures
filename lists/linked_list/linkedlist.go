package linked_list

import (
	"data-structures/utils"
	"fmt"
)

type LinkedList struct {
	dummyHead *item
	size      int
}

type item struct {
	value utils.Equality
	next  *item
}

func (l *LinkedList) String() string {
	res := ""
	for cur := l.dummyHead.next; cur != nil; cur = cur.next {
		res += fmt.Sprintf("%v->", cur.value)
	}
	res += fmt.Sprintf("nil")
	return res
}

func NewLinkedList() *LinkedList {
	return &LinkedList{dummyHead: &item{}}
}

func (l *LinkedList) Add(e utils.Equality) {
	prev := l.dummyHead
	for ; prev.next != nil; prev = prev.next {
	}
	prev.next = &item{value: e, next: prev.next}
	l.size++
}

func (l *LinkedList) Insert(idx int, e utils.Equality) bool {
	if !l.withinIndex(idx) {
		return false
	}

	prev := l.dummyHead
	for i := 0; i < idx; i++ {
		prev = prev.next
	}

	prev.next = &item{e, prev.next}
	l.size++
	return true
}

func (l *LinkedList) Remove(idx int) (utils.Equality, bool) {
	if !l.withinIndex(idx) {
		return nil, false
	}

	prev := l.dummyHead
	for i := 0; i < idx; i++ {
		prev = prev.next
	}

	del := prev.next
	prev.next = del.next
	del.next = nil
	l.size--
	return del.value, true
}

func (l *LinkedList) Set(idx int, e utils.Equality) {
	if !l.withinIndex(idx) {
		return
	}

	cur := l.dummyHead.next
	for i := 0; i < idx; i++ {
		cur = cur.next
	}
	cur.value = e
}

func (l *LinkedList) Get(idx int) (utils.Equality, bool) {
	if !l.withinIndex(idx) {
		return nil, false
	}

	cur := l.dummyHead.next
	for i := 0; i < idx;i++ {
		cur = cur.next
	}
	return cur.value, true
}

func (l *LinkedList) Find(e utils.Equality) int {
	if e != nil && !l.Empty() {
		for i, cur := 0, l.dummyHead.next; cur.next != nil; i, cur = i+1, cur.next {
			if cur.value.Equals(e) {
				return i
			}
		}
	}
	return -1
}

func (l *LinkedList) Contains(e utils.Equality) bool {
	return l.Find(e) != -1
}

func (l *LinkedList) Empty() bool {
	return l.size == 0
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Clear() {
	l.size = 0
	l.dummyHead.next = nil
}

func (l *LinkedList) withinIndex(idx int) bool {
	return idx >= 0 && idx < l.size
}
