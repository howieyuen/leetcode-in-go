package _284_peeking_iterator

// Below is the interface for Iterator, which is already defined for you.

type Iterator struct {
}

func (iter *Iterator) hasNext() bool {
	return false
}

func (iter *Iterator) next() int {
	return -1
}

type PeekingIterator struct {
	iterator    *Iterator
	hasPeeked   bool
	peekElement interface{}
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iterator:    iter,
		hasPeeked:   false,
		peekElement: nil,
	}
}

func (peekIter *PeekingIterator) hasNext() bool {
	return peekIter.hasPeeked || peekIter.iterator.hasNext()
}

func (peekIter *PeekingIterator) next() int {
	if !peekIter.hasPeeked {
		return peekIter.iterator.next()
	}
	res := peekIter.peekElement
	peekIter.hasPeeked = false
	peekIter.peekElement = nil
	return res.(int)
}

func (peekIter *PeekingIterator) peek() int {
	if !peekIter.hasPeeked {
		peekIter.peekElement = peekIter.iterator.next()
		peekIter.hasPeeked = true
	}
	return peekIter.peekElement.(int)
}
