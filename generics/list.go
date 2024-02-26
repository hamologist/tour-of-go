package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) add(n T) {
	l.next = &List[T]{
		next: nil,
		val:  n,
	}
}

func (l *List[T]) toSlice() []T {
	var result []T

	for n := l; n != nil; n = n.next {
		result = append(result, n.val)
	}

	return result
}

func main() {
	myList := List[int]{
		val: 100,
	}

	next := &myList
	for i := 200; i <= 1000; i += 100 {
		next.add(i)
		next = next.next
	}

	fmt.Println(myList.toSlice())
}
