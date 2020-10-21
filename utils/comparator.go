package utils

type Comparator interface {
	Less(than interface{}) bool
	Equality
}

type ComparatorArray []Comparator

func (ca ComparatorArray) Len() int {
	return len(ca)
}

func (ca ComparatorArray) Less(i, j int) bool {
	return ca[i].Less(ca[j])
}

func (ca ComparatorArray) Swap(i, j int) {
	ca[i], ca[j] = ca[j], ca[i]
}