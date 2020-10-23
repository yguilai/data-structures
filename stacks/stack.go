package stacks

import "github.com/yguilai/data-structures/utils"

type Stack interface {
	Push(val utils.Equality)
	Pop() (val utils.Equality, ok bool)
	Peek() (val utils.Equality, ok bool)
	utils.Container
}
