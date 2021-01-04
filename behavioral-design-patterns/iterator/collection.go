package iterator

type Collection interface {
	createIterator() iterator
}
