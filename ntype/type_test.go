package ntype

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Enqueue([]byte(strconv.Itoa(i)))
	}
	fmt.Println("Dequeue: ", string(*q.Dequeue()))
	for _, item := range q.items {
		fmt.Println(string(item))
	}
}

func TestStack(t *testing.T) {
	s := NewStack()
	for i := 0; i < 10; i++ {
		s.Push([]byte(strconv.Itoa(i)))
	}
	fmt.Println("Pop: ", string(*s.Pop()))
	for _, item := range s.items {
		fmt.Println(string(item))
	}
}
