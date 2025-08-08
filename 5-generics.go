package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// generic types
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (s *List[T]) addVal(val T) {
    s.val = val;
}

func (s *List[T]) addNext(val T) {
	var next = &List[T]{val: val}
	s.next = next
}

func main() {
	// type parameters
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	// trying out generics
	var firstItem = &List[string]{}
	firstItem.addVal("first");
	firstItem.addNext("second")

	fmt.Printf("First item value: %v\n", firstItem.val)
	fmt.Printf("First item next: %v\n", *firstItem.next)
	fmt.Printf("First item next.value: %v\n", firstItem.next.val)
}
