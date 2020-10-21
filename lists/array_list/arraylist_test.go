package array_list

import "testing"

type Int int

func (i Int) Equals(than interface{}) bool {
	return i == than.(Int)
}

func TestNewArrayList(t *testing.T) {
	list := NewArrayList(5)

	list.Add(Int(2))
	list.Add(Int(1))
	list.Add(Int(3))
	list.Add(Int(4))

	list.Remove(4)
	list.Remove(3)
	list.Remove(2)
	list.Remove(1)

}