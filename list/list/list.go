package list

type List[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Add(item T) {
	n := &node[T]{value: item}

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

func (l *List[T]) Get(index int) T {
	n := l.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n.value
}

func (l *List[T]) Remove(index int) {
	n := l.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
	}
}

func (l *List[T]) Filter(f func(T) bool) *List[T] {
	newList := NewList[T]()
	n := l.head
	for n != nil {
		if f(n.value) {
			newList.Add(n.value)
		}
		n = n.next
	}
	return newList
}

func Map[T, U any](l *List[T], f func(T) U) *List[U] {
	newList := NewList[U]()
	n := l.head
	for n != nil {
		newList.Add(f(n.value))
		n = n.next
	}
	return newList
}

func (l *List[T]) Reduce(f func(T, T) T) T {
	n := l.head
	result := n.value
	for n != nil {
		result = f(result, n.value)
		n = n.next
	}
	return result
}