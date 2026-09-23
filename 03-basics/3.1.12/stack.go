package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	last := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return last, true
}

func stackExample() {
	stack := Stack[Person]{}
	stack.Push(Person{Name: "John", Value: 123})
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
