package segment_tree

type SegmentTree struct {
	size int
	tree []interface{}
	merge Merger
}

type Merger func(a, b interface{}) interface{}

// NewSegmentTree constructor of segment tree
func NewSegmentTree(arr []interface{}, merger Merger) *SegmentTree {
	if len(arr) == 0 {
		panic("Array is empty.")
	}

	t := &SegmentTree{len(arr)-1, make([]interface{}, 4*len(arr)), merger}
	t.buildSegmentTree(arr, 0, 0, len(arr)-1)
	return t
}

// buildSegmentTree 构建线段树
func (t *SegmentTree) buildSegmentTree(arr []interface{}, treeIndex, l, r int) {
	if l == r {
		t.tree[treeIndex] = arr[l]
		return
	}

	mid := l + (r-l)/2
	left := leftChild(treeIndex)
	right := rightChild(treeIndex)

	t.buildSegmentTree(arr, left, l, mid)
	t.buildSegmentTree(arr, right, mid+1, r)
	t.tree[treeIndex] = t.merge(t.tree[left], t.tree[right])
}

// leftChild 获取节点左孩子的索引
func leftChild(i int) int {
	return 2*i + 1
}

// rightChild 获取节点右孩子的索引
func rightChild(i int) int {
	return 2*i + 2
}

// Query 返回[ql, qr]的值
func (t *SegmentTree) Query(ql, qr int) interface{} {
	if ql < 0 || ql >= t.size+1 ||
		qr < 0 || qr >= t.size+1 || ql > qr {
		panic("Index is illegal.")
	}

	return t.query(0, 0, t.size, ql, qr)
}

// query 在以treeIndex为根的线段树中[l, r]的范围内, 搜索区间[ql, qr]的值
func (t *SegmentTree) query(treeIndex, l, r, ql, qr int) interface{} {
	if l == ql && r == qr {
		return t.tree[treeIndex]
	}

	mid := l + (r-l)/2
	left := leftChild(treeIndex)
	right := rightChild(treeIndex)

	if ql >= mid+1 {
		return t.query(right, mid+1, r, ql, qr)
	} else if qr <= mid {
		return t.query(left, l, mid, ql, qr)
	}

	return t.merge(t.query(left, l, mid, ql, mid), t.query(right, mid+1, r, mid+1, qr))
}

// Set 设置索引为idx的节点值为e
func (t *SegmentTree) Set(idx int, e interface{}) {
	if idx < 0 || idx >= t.size+1 {
		panic("Index is illegal")
	}

	t.set(0, 0, t.size, idx, e)
}

func (t *SegmentTree) set(treeIdx, l, r, idx int, e interface{}) {
	if l == r {
		t.tree[treeIdx] = e
		return
	}

	mid := l + (r-l)/2
	left := leftChild(treeIdx)
	right := rightChild(treeIdx)
	if idx >= mid+1 {
		t.set(right, mid+1, r, idx, e)
	} else {
		t.set(left, l, mid, idx, e)
	}

	// idx索引在treeIdx索引节点的子节点, 所以需要更新treeIdx所在节点的值
	t.tree[treeIdx] = t.merge(t.tree[left], t.tree[right])
}

func (t *SegmentTree) Empty() bool {
	return len(t.tree) == 0
}

func (t *SegmentTree) Size() int {
	return t.size
}

func (t *SegmentTree) Clear() {
	t.tree = make([]interface{}, 0)
}

