package main

import (
	"errors"
	"sync"
)

type stack struct {
	lock   sync.Mutex // you don't have to do this if you don't want thread safety
	s      []Node
	lenght int
}

func Stack() *stack {
	return &stack{sync.Mutex{}, make([]Node, 0), 0}
}

func (s *stack) Push(v Node) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
	s.lenght++
}

func (s *stack) Pop() (Node, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := s.lenght
	if l == 0 {
		return Node{}, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	s.lenght--
	return res, nil
}

func (s *stack) Length() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.lenght
}
