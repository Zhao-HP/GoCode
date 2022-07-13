package main

import "fmt"

type Result[T any] struct {
	Value T
}

func main() {
	r := &Result[int]{
		Value: 5,
	}

	fmt.Println(r)
}

func (n Result[T]) getValue() T {
	return n.Value
}
