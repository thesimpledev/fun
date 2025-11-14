// Package stack holds our stack
package stack

import "sync"

type Stack struct {
	items []int
	mu    sync.RWMutex
}

func New() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.items) == 0 {
		return 0, false
	}
	pos := len(s.items) - 1
	v := s.items[pos]
	s.items = s.items[:pos]
	return v, true
}

func (s *Stack) Peek() (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.items) == 0 {
		return 0, false
	}

	pos := len(s.items) - 1

	return s.items[pos], true
}

func (s *Stack) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}
