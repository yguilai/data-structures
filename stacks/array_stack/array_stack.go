package array_stack

import (
	"data-structures/lists/array_list"
	"data-structures/utils"
)

type ArrayStack struct {
	list *array_list.ArrayList
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{list: array_list.NewArrayList(10)}
}

func NewArrayStackWithCap(c int) *ArrayStack {
	return &ArrayStack{list: array_list.NewArrayList(c)}
}

func (a *ArrayStack) Push(val utils.Equality) {
	a.list.Add(val)
}

func (a *ArrayStack) Pop() (val utils.Equality, ok bool) {
	return a.list.Remove(a.list.Size() - 1)
}

func (a *ArrayStack) Peek() (val utils.Equality, ok bool) {
	return a.list.Get(a.list.Size() - 1)
}

func (a *ArrayStack) Empty() bool {
	return a.list.Empty()
}

func (a *ArrayStack) Size() int {
	return a.list.Size()
}

func (a *ArrayStack) Clear() {
	a.list.Clear()
}
