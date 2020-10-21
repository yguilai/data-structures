package avl_tree

import (
	"data-structures/utils"
	"math"
)

type AVL struct {
	Root *Node
	size uint
}

func NewAVL() *AVL {
	return &AVL{}
}

func (avl *AVL) Empty() bool {
	return avl.size == 0
}

func (avl *AVL) Size() int {
	return int(avl.size)
}

func (avl *AVL) Clear() {
	avl.size = 0
	avl.Root = nil
}

func (avl *AVL) Put(key utils.Comparator, val interface{}) {
	if key == nil {
		return
	}

	avl.Root = avl.Root.put(key, val, &avl.size)
}

func (avl *AVL) Has(key utils.Comparator) bool {
	_, f := avl.Get(key)
	return f
}

func (avl *AVL) Get(key utils.Comparator) (val interface{}, found bool) {
	return avl.Root.get(key)
}

func (avl *AVL) Remove(key utils.Comparator) bool {
	var removed bool
	avl.Root, removed = avl.Root.remove(key, &avl.size)
	return removed
}

// -----------------------

type Node struct {
	Left, Right *Node
	Key         utils.Comparator
	Value       interface{}
	Height      int
}

func NewNode(key utils.Comparator, val interface{}) *Node {
	return &Node{Key: key, Value: val, Height: 1}
}

func (n *Node) put(key utils.Comparator, val interface{}, size *uint) *Node {
	if n == nil {
		*size++
		return NewNode(key, val)
	}

	if key.Less(n.Key) {
		n.Left = n.Left.put(key, val, size)
	} else if n.Key.Less(key) {
		n.Right = n.Right.put(key, val, size)
	} else {
		n.Value = val
	}

	n.Height = 1 + int(math.Max(getHeight(n.Left), getHeight(n.Right)))
	return n.balance()
}

func (n *Node) balance() *Node {
	bf := getBalanceFactor(n)

	// LL
	if bf > 1 && getBalanceFactor(n.Left) >= 0 {
		return n.rightRotate()
	}
	// RR
	if bf < -1 && getBalanceFactor(n.Right) <= 0 {
		return n.leftRotate()
	}
	// LR
	if bf > 1 && getBalanceFactor(n.Left) <= 0 {
		n.Left = n.Left.leftRotate()
		return n.rightRotate()
	}
	// RL
	if bf < -1 && getBalanceFactor(n.Right) >= 0 {
		n.Right = n.Right.rightRotate()
		return n.leftRotate()
	}

	return n
}

// rightRotate 对当前节点进行右旋操作(LL型)
// 	   	 n                       x
//      / \					   /   \
//     x  T4     右旋(n)      z     n
//    / \		------->    / \    / \
//   z  T3  			   T1 T2  T3 T4
//  / \
// T1 T2
func (n *Node) rightRotate() *Node {
	x, T3 := n.Left, n.Left.Right
	x.Right, n.Left = n, T3

	x.Height = 1 + int(math.Max(getHeight(x.Left), getHeight(x.Right)))
	n.Height = 1 + int(math.Max(getHeight(n.Left), getHeight(n.Right)))
	return x
}

// leftRotate 对当前节点进行左旋操作 (RR型)
//   n					       x
//  / \                      /   \
// T4  x        左旋(n)     n     z
//    / \      ------->   / \    / \
//   T3  z               T4 T3  T2 T1
//      / \
//     T1 T2
func (n *Node) leftRotate() *Node {
	x, T3 := n.Right, n.Right.Left
	x.Left, n.Right = n, T3

	x.Height = 1 + int(math.Max(getHeight(x.Left), getHeight(x.Right)))
	n.Height = 1 + int(math.Max(getHeight(n.Left), getHeight(n.Right)))
	return x
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

	var (
		removed bool
		resNode *Node
	)
	if key.Less(n.Key) {
		n.Left, removed = n.Left.remove(key, size)
		resNode = n
	} else if n.Key.Less(key) {
		n.Right, removed = n.Right.remove(key, size)
		resNode = n
	} else {
		if n.Left == nil {
			resNode = n.Right
			n.Right = nil
		} else if n.Right == nil {
			resNode = n.Left
			n.Left = nil
		} else {
			successor := n.Right.min()
			successor.Right, _ = n.Right.remove(successor.Key, size)
			successor.Left = n.Left
			n.Left, n.Right = nil, nil
			resNode = successor
		}
	}

	if !removed {
		return resNode, false
	}

	// 若删除的是叶子节点 resNode可能为空
	if resNode == nil {
		return nil, true
	}

	resNode.Height = 1 + int(math.Max(getHeight(resNode.Left), getHeight(resNode.Right)))
	return resNode.balance(), true
}

// min 获取最小节点
func (n *Node) min() *Node {
	if n.Left == nil {
		return n
	}

	return n.Left.min()
}

// getHeight 获取节点高度
func getHeight(n *Node) float64 {
	if n == nil {
		return 0
	}
	return float64(n.Height)
}

// getBalanceFactor 计算节点平衡因子
func getBalanceFactor(n *Node) float64 {
	if n == nil {
		return 0
	}
	return getHeight(n.Left) - getHeight(n.Right)
}