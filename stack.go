package main

import (
	"errors"
	"sync"
)

type stack struct {
	lock   sync.Mutex // you don't have to do this if you don't want thread safety
	s      []XMLElement
	lenght int
}

func Stack() *stack {
	return &stack{sync.Mutex{}, make([]XMLElement, 0), 0}
}

func (s *stack) Push(v XMLElement) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.push(v)
}

func (s *stack) Pop() (XMLElement, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.pop()
}

func (s *stack) push(v XMLElement) {
	s.s = append(s.s, v)
	s.lenght++
}

func (s *stack) pop() (XMLElement, error) {
	l := s.lenght
	if l == 0 {
		return XMLElement{}, errors.New("Empty Stack")
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

func (s *stack) UpdateValueLastElemet(in string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	node, err := s.pop()
	if err != nil {
		return err
	}

	node.InnerValue = in
	s.push(node)
	return nil
}
