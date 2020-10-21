package trie

type Trie struct {
	root *node
	size int
}

func (t *Trie) Empty() bool {
	return t.size == 0
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) Clear() {
	t.root = nil
}

type node struct {
	last bool
	next map[rune]*node
}

func newNode() *node {
	return &node{last: false, next: make(map[rune]*node)}
}

// NewTrie constructor of Trie
func NewTrie() *Trie {
	return &Trie{newNode(), 0}
}

// Add 向trie中添加一个新的单词word
func (t *Trie) Add(word string) {
	cur := t.root
	for _, v := range word {
		if _, ok := cur.next[v]; !ok {
			cur.next[v] = newNode()
		}
		cur = cur.next[v]
	}

	// 如果新添加的单词结尾字符是已存在某个单词的部分, 对这个字符添加结束标识
	if !cur.last{
		cur.last = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for _, v := range word {
		if _, ok := cur.next[v]; !ok {
			return false
		}
		cur = cur.next[v]
	}
	return cur.last
}

// IsPrefix 查询Trie中是否有单词以prefix为前缀
func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root
	for _, v := range prefix {
		if _, ok := cur.next[v]; !ok {
			return false
		}
		cur = cur.next[v]
	}
	return true
}

// Remove 递归实现Trie删除操作
func (t *Trie) Remove(word string) {
	t.remove(t.root, word, 0)
}

func (t *Trie) remove(n *node, word string, idx int) bool {
	if idx == len(word) && n.last {
		if len(n.next) == 0 {
			n.next = map[rune]*node{}
			return true
		}
		return false
	}

	c := rune(word[idx])
	if _, ok := n.next[c]; ok {
		if t.remove(n.next[c], word, idx+1) {
			if !n.last && len(n.next) == 1 {
				n.next = map[rune]*node{}
				return true
			}
			return false
		}
	}
	return false
}