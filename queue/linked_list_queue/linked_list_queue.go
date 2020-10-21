package linked_list_queue

import (
	"data-structures/lists/linked_list"
	"data-structures/utils"
)

type ArrayQueue struct {
	list *linked_list.LinkedList
}

func (q *ArrayQueue) Enqueue(e utils.Equality) {
	q.list.Add(e)
}

func (q *ArrayQueue) Dequeue() (e utils.Equality, ok bool) {
	return q.list.Remove(0)
}

func (q *ArrayQueue) GetFront() (e utils.Equality, ok bool) {
	return q.list.Get(0)
}

func (q *ArrayQueue) Empty() bool {
	return q.list.Empty()
}

func (q *ArrayQueue) Size() int {
	return q.list.Size()
}

func (q *ArrayQueue) Clear() {
	q.list.Clear()
}

