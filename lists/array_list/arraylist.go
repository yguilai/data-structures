package array_list

import (
	"data-structures/utils"
)

type ArrayList struct {
	data []utils.Equality
	size int
}

func NewArrayList(cap int) *ArrayList {
	return &ArrayList{data: make([]utils.Equality, cap)}
}

func (al *ArrayList) Add(c utils.Equality) {
	al.data[al.size] = c
	al.size++
	al.expand()
}

func (al *ArrayList) Insert(idx int, e utils.Equality) bool {
	if !al.withinIndex(idx) {
		return false
	}

	for i := al.size; i > idx; i-- {
		al.data[i+1] = al.data[i]
	}

	al.data[idx] = e
	al.size++
	al.expand()
	return true
}

func (al *ArrayList) Remove(idx int) (utils.Equality, bool) {
	if !al.withinIndex(idx) {
		return nil, false
	}

	res := al.data[idx]
	for i := idx + 1; i < al.size; i++ {
		al.data[i-1] = al.data[i]
	}
	al.size--
	al.data[al.size] = nil

	al.shrink()
	return res, true
}

func (al *ArrayList) Set(idx int, e utils.Equality) {
	if !al.withinIndex(idx) {
		return
	}

	al.data[idx] = e
}

func (al *ArrayList) Get(idx int) (utils.Equality, bool) {
	if !al.withinIndex(idx) {
		return nil, false
	}

	return al.data[idx], true
}

func (al *ArrayList) Find(e utils.Equality) int {
	if e != nil && al.size != 0 {
		for i, v := range al.data {
			if v.Equals(e) {
				return i
			}
		}
	}

	return -1
}

func (al *ArrayList) Contains(e utils.Equality) bool {
	return al.Find(e) != -1
}

func (al *ArrayList) withinIndex(idx int) bool {
	return idx >= 0 && idx < al.size
}

func (al *ArrayList) Empty() bool {
	return al.size == 0
}

func (al *ArrayList) Size() int {
	return al.size
}

func (al *ArrayList) Clear() {
	al.size = 0
	al.data = make([]utils.Equality, cap(al.data))
}

func (al *ArrayList) resize(c int) {
	newData := make([]utils.Equality, c, c)
	copy(newData, al.data)
	al.data = newData
}

func (al *ArrayList) expand()  {
	c := cap(al.data)

	if int(float32(al.size)*1.5) >= c {
		al.resize(2*c)
	}

}

func (al *ArrayList) shrink()  {
	c := cap(al.data)
	if al.size <= int(float32(c)*0.25) {
		al.resize(al.size)

	}
}
