package utils

type Container interface {
	Empty() bool
	Size() int
	Clear()
}
