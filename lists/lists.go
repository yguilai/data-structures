package lists

import "data-structures/utils"

type List interface {
	Add(e utils.Equality)
	Insert(idx int, e utils.Equality) bool
	Remove(idx int) (utils.Equality, bool)
	Set(idx int, e utils.Equality)
	Get(idx int) (utils.Equality, bool)
	Find(e utils.Equality) int
	Contains(e utils.Equality) bool
	utils.Container
}
