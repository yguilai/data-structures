package linked_list_stack

import (
	"github.com/yguilai/data-structures/lists/linked_list"
	"github.com/yguilai/data-structures/utils"
)

type LinkedListStack struct {
	list *linked_list.LinkedList
}

func (s *LinkedListStack) Push(val utils.Equality) {
	s.list.Add(val)
}

func (s *LinkedListStack) Pop() (val utils.Equality, ok bool) {
	return s.list.Remove(s.list.Size()-1)
}

func (s *LinkedListStack) Peek() (val utils.Equality, ok bool) {
	return s.list.Get(s.list.Size() -1)
}

func (s *LinkedListStack) Empty() bool {
	return s.list.Empty()
}

func (s *LinkedListStack) Size() int {
	return s.list.Size()
}

func (s *LinkedListStack) Clear() {
	s.list.Clear()
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{list: linked_list.NewLinkedList()}
}


