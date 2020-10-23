package red_black_tree

import (
	"github.com/yguilai/data-structures/utils"
)

// 红黑树性质
// 1. 每个节点或是黑色, 或是红色
// 2. 根节点是黑色
// 3. 每个叶子节点是黑色(指为nil的的叶子节点)
// 4. 如果一个节点是红色, 则它的子节点必须是黑色
// 5. 从一个节点到该节点的子孙节点的所有路径上包含相同数目的黑节点

type (
	LLRBTree struct {
		Root *LLRBTNode
		size uint
	}

	LLRBTNode struct {
		Left, Right *LLRBTNode
		Key         utils.Comparator
		Value       interface{}
		Color       bool
	}
)

const (
	RED   = true
	BLACK = false
)

func NewLLRBTree() *LLRBTree {
	return &LLRBTree{}
}

func isRed(n *LLRBTNode) bool {
	if n == nil {
		return BLACK
	}

	return n.Color
}

// leftRotate 左旋转
//    n						       x
//   / \         左旋转           / \
//  T1  x      --------->       n   T3
//     / \                     / \
//    T2 T3                   T1 T2
func leftRotate(n *LLRBTNode) *LLRBTNode {
	if n == nil {
		return nil
	}

	x := n.Right
	n.Right, x.Left = x.Left, n

	x.Color, n.Color = n.Color, RED
	return x
}

// rightRotate 右旋转
//     n					x
//    / \     右旋转       / \
//   x  T2   -------->   y   n
//  / \						/ \
// y  T1                   T1 T2
func rightRotate(n *LLRBTNode) *LLRBTNode {
	if n == nil {
		return nil
	}

	x := n.Left
	n.Left, x.Right = x.Right, n
	x.Color, n.Color = n.Color, RED
	return x
}

// flipColors 颜色反转
func flipColors(n *LLRBTNode) {
	if n == nil {
		return
	}
	n.Color = !n.Color
	n.Left.Color = !n.Left.Color
	n.Right.Color = !n.Left.Color
}

// moveRedToLeft 红色左移
func moveRedToLeft(n *LLRBTNode) *LLRBTNode {
	flipColors(n)

	if isRed(n.Right.Left) {
		n.Right = rightRotate(n.Right)
		n = leftRotate(n)
		flipColors(n)
	}
	return n
}

// moveRedToRight 红色右移
func moveRedToRight(n *LLRBTNode) *LLRBTNode {
	flipColors(n)

	if isRed(n.Left.Left) {
		n = rightRotate(n)
		flipColors(n)
	}

	return n
}

// Put 添加/修改元素
func (tree *LLRBTree) Put(key utils.Comparator, val interface{}) {
	tree.Root = tree.Root.put(key, val, &tree.size)
}

func (n *LLRBTNode) put(key utils.Comparator, val interface{}, size *uint) *LLRBTNode {
	if n == nil {
		*size++
		return &LLRBTNode{Key: key, Value: val, Color: RED}
	}

	if key.Less(n.Key) {
		n.Left = n.Left.put(key, val, size)
	} else if n.Key.Less(key) {
		n.Right = n.Right.put(key, val, size)
	} else {
		n.Value = val
	}

	node := n

	if isRed(node.Right) && !isRed(node.Left) {
		node = leftRotate(node)
	} else {
		if isRed(node.Left) && isRed(node.Left.Left) {
			node = rightRotate(node)
		}

		if isRed(node.Left) && isRed(node.Right) {
			flipColors(node)
		}
	}

	return node
}

// Get 获取
func (tree *LLRBTree) Get(key utils.Comparator) (interface{}, bool) {
	return tree.Root.get(key)
}

func (n *LLRBTNode) get(key utils.Comparator) (interface{}, bool) {
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

func (tree *LLRBTree) Has(key utils.Comparator) bool {
	_, has := tree.Get(key)
	return has
}

func (tree *LLRBTree) Delete(key utils.Comparator) interface{} {
	if _, has := tree.Get(key); !has {
		return nil
	}

	var (
		r       = tree.Root
		deleted interface{}
	)

	if !isRed(r.Left) && !isRed(r.Right) {
		r.Color = RED
	}

	r, deleted = r.delete(key)
	if r != nil {
		r.Color = BLACK
	}
	return deleted
}

func (n *LLRBTNode) delete(key utils.Comparator) (*LLRBTNode, interface{}) {
	var (
		node    = n
		deleted interface{}
	)
	if key.Less(node.Key) {
		if !isRed(node.Left) && !isRed(node.Left.Left) {
			node = moveRedToLeft(node)
		}

		node.Left, deleted = node.Left.delete(key)
	} else {
		if isRed(node.Left) {
			node = rightRotate(node)
		}

		if key.Equals(node.Key) && node.Right == nil {
			return nil, node.Value
		}

		if !isRed(node.Right) && !isRed(node.Right.Left) {
			node = moveRedToRight(node)
		}

		if key.Equals(node.Key) {
			successor := node.Right.min()
			node.Value = successor.Value
			node.Key = successor.Key

			node.Right = node.Right.deleteMin()
		} else {
			node.Right, deleted = node.Right.delete(key)
		}
	}

	return node.fixUp(), deleted
}

func (n *LLRBTNode) min() *LLRBTNode {
	if n.Left == nil {
		return n
	}

	return n.Left.min()
}

func (n *LLRBTNode) deleteMin() *LLRBTNode {
	node := n

	if node.Left == nil {
		return nil
	}

	if !isRed(node.Left) && !isRed(node.Left.Left) {
		node = moveRedToLeft(node)
	}

	node.Left = node.Left.deleteMin()
	return node.fixUp()
}

func (n *LLRBTNode) fixUp() *LLRBTNode {
	node := n
	if isRed(node.Right) {
		node = leftRotate(node)
	}

	if isRed(node.Left) && isRed(node.Left.Left) {
		node = rightRotate(node)
	}

	if isRed(node.Left) && isRed(node.Right) {
		flipColors(node)
	}

	return node
}
