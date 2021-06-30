package stack

import (
    "errors"
)

type (
    Stack struct {
        top *Node
        size uint
    }

    Node struct {
        value uint
        prev *Node
    }
)

func New() *Stack {
    return &Stack{nil, 0}
}

func (s *Stack) Pop() (uint, error) {
    if s.top == nil {
        return 0, errors.New("empty stack")
    }

    node := s.top
    s.top = node.prev
    s.size--

    return node.value, nil
}

func (s *Stack) Push(value uint) {
    node := &Node{value,s.top}
    s.top = node
    s.size++
}

func (s *Stack) Size() uint {
    return s.size
}
