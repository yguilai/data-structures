package binary_heap

import (
	"data-structures/utils"
)

type MaxHeap struct {
	data utils.ComparatorArray
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data: make(utils.ComparatorArray, 0)}
}

func NewMaxHeapWithHeapify(arr utils.ComparatorArray) *MaxHeap {
	m := &MaxHeap{arr}
	for i := parent(m.data.Len() - 1); i >= 0; i-- {
		m.siftDown(i)
	}

	return m
}

func (m *MaxHeap) Add(e utils.Comparator) {
	m.data = append(m.data, e)
	m.siftUp(m.data.Len() - 1)
}


func (m *MaxHeap) ExtractMax() utils.Comparator {
	r := m.max()
	m.data.Swap(0, m.data.Len()-1)
	m.data = m.data[:m.data.Len()-1]
	m.siftDown(0)
	return r
}

func parent(index int) int {
	if index == 0 {
		panic("index-0 doesn't have parent.")
	}
	return (index - 1) / 2
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

// siftUp 元素上浮
func (m *MaxHeap) siftUp(i int) {
	for ; i > 0 && m.data.Less(parent(i), i); i = parent(i) {
		m.data.Swap(parent(i), i)
	}
}

// siftDown 元素下沉
func (m *MaxHeap) siftDown(i int) {
	for leftChild(i) < m.data.Len() {
		j := leftChild(i)

		if j+1 < m.data.Len() && m.data.Less(j, j+1) {
			// 确保m.data[j] 是leftChild和rightChild中的最大值
			j = rightChild(i)
		}

		if m.data.Less(j, i) {
			break
		}

		m.data.Swap(i, j)
		i = j
	}
}

func (m *MaxHeap) max() utils.Comparator {
	if m.data.Len() == 0 {
		return nil
	}

	return m.data[0]
}


func (m *MaxHeap) Replace(new utils.Comparator) utils.Comparator {
	r := m.max()
	m.data[0] = new
	m.siftDown(0)
	return r
}

