package binary_search_tree

import (
	"github.com/yguilai/data-structures/utils"
)

type BST struct {
	Root *Node
	size uint
}

func (tree *BST) Empty() bool {
	return tree.size == 0
}

func (tree *BST) Size() int {
	return int(tree.size)
}

func (tree *BST) Clear() {
	tree.Root = nil
}

func (tree *BST) Put(key utils.Comparator, val interface{}) {
	if key == nil {
		return
	}

	tree.Root = tree.Root.put(key, val, &tree.size)
}

func (tree *BST) Get(key utils.Comparator) (val interface{}, found bool) {
	tree.Root.get(key)
	return nil, false
}

func (tree *BST) Has(key utils.Comparator) bool {
	_, found := tree.Root.get(key)
	return found
}

func (tree *BST) Remove(key utils.Comparator) bool {
	var deleted bool
	tree.Root, deleted = tree.Root.remove(key, &tree.size)
	return deleted
}

// -----------

type Node struct {
	Left, Right *Node
	Key         utils.Comparator
	Value       interface{}
}

func (n *Node) put(key utils.Comparator, val interface{}, size *uint) *Node {
	if n == nil {
		*size++
		return &Node{Key: key, Value: val}
	}

	if key.Less(n.Key) {
		n.Left = n.Left.put(key, val, size)
	} else if n.Key.Less(key) {
		n.Right = n.Right.put(key, val, size)
	}

	n.Value = val
	return n
}

func (n *Node) get(key utils.Comparator) (val interface{}, found bool) {
	if n == nil {
		return nil, false
	}

	if key.Less(n.Key) {
		return n.Left.get(key)
	} else if n.Key.Less(key) {
		return n.Right.get(key)
	}

	return n.Value, true
}

func (n *Node) remove(key utils.Comparator, size *uint) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	var deleted bool
	if key.Less(n.Key) {
		n.Left, deleted = n.Left.remove(key, size)
		return n, deleted
	} else if n.Key.Less(key) {
		n.Right, deleted = n.Right.remove(key, size)
		return n, deleted
	}

	if n.Left == nil {
		r := n.Right
		n.Right = nil
		*size--
		return r, true
	}

	if n.Right == nil {
		l := n.Left
		n.Left = nil
		*size--
		return l, true
	}

	successor := n.Right.min()
	successor.Left = n.Left
	successor.Right, _ = n.Right.remove(successor.Key, size)
	n.Left, n.Right = nil, nil
	return successor, true
}

func (n *Node) min() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.min()
}
