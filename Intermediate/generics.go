package main

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	// x, y := 1, 2
	// x, y = swap(x, y)
	// println(x, y)

	intStack := Stack[int]{}
	intStack.push(10)
	intStack.push(20)
	val, ok := intStack.pop()
	if ok {
		println(val)
	}

	stringStack := Stack[string]{}
	stringStack.push("hello")
	stringStack.push("world")
	strVal, ok := stringStack.pop()
	if ok {
		println(strVal)
	}

}
