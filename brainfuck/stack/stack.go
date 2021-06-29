package stack

import (
    "errors"
)

type (
    Stack struct {
        top *Node
    }

    Node struct {
        value uint
        prev *Node
    }
)

func New() *Stack {
    return &Stack{nil}
}

func (s *Stack) Pop() (uint, error) {
    if s.top == nil {
        return 0, errors.New("empty stack")
    }

    node := s.top
    s.top = node.prev

    return node.value, nil
}

func (s *Stack) Push(value uint) {
    node := &Node{value,s.top}
    s.top = node
}

