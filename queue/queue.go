package queue

import "data-structures/utils"

type Queue interface {
	Enqueue(e utils.Equality)
	Dequeue() (e utils.Equality, ok bool)
	GetFront() (e utils.Equality, ok bool)
	utils.Container
}
