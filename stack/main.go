package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Length() int {
	return len(s.data)
}

func (s *Stack[T]) Pop() (val T, ok bool) {
	length := len(s.data)

	if length == 0 {
		var zero T
		return zero, false
	}
	lastIndex := length - 1
	val = s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return val, true
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *Stack[T]) Peek() (val T, ok bool) {
	length := len(s.data)

	if length == 0 {
		var zero T
		return zero, false
	}

	return s.data[length-1], true
}

func do1() error {
	// log.Fatal("fatal do1")
	return errors.New("error do1")
}

func do2() error {
	err := do1()
	if err != nil {
		return fmt.Errorf("error do2 %s", err)
	}

	return fmt.Errorf("default error")
}

func main() {
	// stack := New[int]()

	// stack.Push(0)
	// stack.Push(1)

	// peeked, _ := stack.Peek()

	// fmt.Println(stack, peeked)
	err := do2().Error()
	fmt.Println(err)

}
