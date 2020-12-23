package hashtable

import "errors"

type Hashabler interface {
	Hash() int
	Equals(than Hashabler) bool
}

type HashTable struct {
	table      []*HashItem
	size       int
	primeIndex int
	upper      int
	lower      int
}

var primes = []int{
	53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157,
	98317, 196613, 393241, 786433, 1572869, 3145739, 6291469,
	12582917, 25165843, 50331653, 100663319, 201326611,
	402653189, 805306457, 1610612741,
}

func NewHashTable() *HashTable {
	return &HashTable{
		table:      make([]*HashItem, primes[0]),
		size:       0,
		primeIndex: 0,
		upper:      10,
		lower:      2,
	}
}

func (t *HashTable) Size() int {
	return t.size
}

func (t *HashTable) hash(key Hashabler) int {
	return key.Hash() % primes[t.primeIndex]
}

func (t *HashTable) SetUpperAndLower(u, l int) error {
	t.upper = u
	t.lower = l
	if t.size > t.upper*primes[t.primeIndex] && t.primeIndex+1 <= len(primes) {
		t.primeIndex++
		return t.resize()
	}

	if t.size < t.lower*primes[t.primeIndex] && t.primeIndex-1 >= 0 {
		t.primeIndex--
		return t.resize()
	}

	return nil
}

func (t *HashTable) Put(key Hashabler, val interface{}) error {
	h := t.hash(key)
	var appended bool
	t.table[h], appended = t.table[h].Put(key, val)
	if appended {
		t.size++
	}

	if t.size > t.upper*primes[t.primeIndex] && t.primeIndex+1 < len(primes) {
		t.primeIndex++
		return t.resize()
	}
	return nil
}

func (t *HashTable) Get(key Hashabler) (val interface{}, err error) {
	if val, has := t.table[t.hash(key)].Get(key); has {
		return val, nil
	}
	return nil, errors.New("Key is not found.")
}

func (t *HashTable) Hash(key Hashabler) bool {
	_, has := t.table[t.hash(key)].Get(key)
	return has
}

// Remove remove item by the key that implement Hashabler interface
func (t *HashTable) Remove(key Hashabler) (val interface{}, err error) {
	l := t.table[t.hash(key)]
	val, has := l.Get(key)
	if !has {
		return nil, errors.New("Key is not found.")
	}
	l = l.Remove(key)
	t.size--
	if t.size < t.lower*primes[t.primeIndex] && t.primeIndex-1 >= 0 {
		t.primeIndex--
		return val, t.resize()
	}
	return val, nil
}

func (t *HashTable) resize() error {
	table := t.table
	t.table = make([]*HashItem, primes[t.primeIndex])
	t.size = 0
	for _, item := range table {
		for i := item; i != nil; i = i.Next {
			if err := t.Put(i.Key, i.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

// ------
type HashItem struct {
	Key   Hashabler
	Value interface{}
	Next  *HashItem
}

func (t *HashItem) Get(key Hashabler) (val interface{}, has bool) {
	if t == nil {
		return nil, false
	} else if t.Key.Equals(key) {
		return t.Value, true
	} else {
		return t.Next.Get(key)
	}
}

func (t *HashItem) Put(key Hashabler, val interface{}) (item *HashItem, appended bool) {
	if t == nil {
		return &HashItem{key, val, nil}, true
	} else if t.Key.Equals(key) {
		t.Value = val
		return t, false
	} else {
		t.Next, appended = t.Next.Put(key, val)
		return t, appended
	}
}

func (t *HashItem) Remove(key Hashabler) *HashItem {
	if t == nil {
		panic("Key is not found.")
	}

	if t.Key.Equals(key) {
		return t.Next
	} else {
		t.Next = t.Next.Remove(key)
		return t
	}
}