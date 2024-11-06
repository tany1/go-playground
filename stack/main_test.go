package main

import "testing"

func TestLength(t *testing.T) {
	intStack := New[int]()
	stringStack := New[string]()

	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	if intStack.Length() != 3 {
		t.Errorf("got %d, expected %d", intStack.Length(), 3)
	}

	if stringStack.Length() != 0 {
		t.Errorf("got %d, expected %d", intStack.Length(), 0)
	}

	stringStack.Push("abc")
	if stringStack.Length() != 1 {
		t.Errorf("got %d, expected %d", intStack.Length(), 1)
	}

	stringStack.Push("def")
	if stringStack.Length() != 2 {
		t.Errorf("got %d, expected %d", intStack.Length(), 2)
	}
}

func TestPeek(t *testing.T) {
	intStack := New[int]()

	_, ok := intStack.Peek()

	if ok {
		t.Errorf("stack should be empty")
	}

	intStack.Push(10)
	intStack.Push(20)

	peek, _ := intStack.Peek()

	if peek != 20 {
		t.Errorf("got %d, expected %d", peek, 20)
	}
}

func TestPush(t *testing.T) {
	intStack := New[int]()

	intStack.Push(10)
	peek, _ := intStack.Peek()
	if peek != 10 {
		t.Errorf("got %d, expected %d", peek, 10)
	}

	intStack.Push(20)
	peek, _ = intStack.Peek()
	if peek != 20 {
		t.Errorf("got %d, expected %d", peek, 20)
	}

	intStack.Push(30)
	peek, _ = intStack.Peek()
	if peek != 30 {
		t.Errorf("got %d, expected %d", peek, 30)
	}
}

func TestPop(t *testing.T) {
	intStack := New[int]()

	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	val, _ := intStack.Pop()
	if val != 3 {
		t.Errorf("got %d, expected %d", val, 3)
	}

	val, _ = intStack.Pop()
	if val != 2 {
		t.Errorf("got %d, expected %d", val, 2)
	}

	val, _ = intStack.Pop()
	if val != 1 {
		t.Errorf("got %d, expected %d", val, 1)
	}

	_, ok := intStack.Pop()
	if ok {
		t.Errorf("stack should be empty")
	}
}
