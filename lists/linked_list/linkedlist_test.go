package linked_list

import (
	"fmt"
	"testing"
)

type Int int

func (i Int) Equals(than interface{}) bool {
	return i == than.(Int)
}

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()

	list.Add(Int(1))
	list.Add(Int(2))
	list.Add(Int(3))

	list.Insert(2, Int(1))

	fmt.Println(list)
	list.Remove(0)
	fmt.Println(list)
	list.Remove(2)
	fmt.Println(list)

	fmt.Println(list.Get(0))
	fmt.Println(list.Find(Int(2)))
	fmt.Println(list.Contains(Int(2)))

}
