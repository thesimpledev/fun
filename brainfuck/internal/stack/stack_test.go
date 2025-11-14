package stack

import "testing"

func TestNew(t *testing.T) {
	s := New()

	if s == nil {
		t.Fatal("stack was nil and should not be")
	}
}

func TestPush(t *testing.T) {
	s := New()
	want := 1
	s.Push(want)
	got, _ := s.Pop()

	if want != got {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPopEmpty(t *testing.T) {
	s := New()
	_, exists := s.Pop()
	if exists {
		t.Error("Exists is True and should not be")
	}
}

func TestPop(t *testing.T) {
	s := New()
	want := 1
	s.Push(5)
	s.Push(want)
	num, exists := s.Pop()
	if !exists {
		t.Errorf("doesn't exist and should")
	}

	if num != want {
		t.Errorf("got %d want %d", num, want)
	}

	if s.Len() != 1 {
		t.Errorf("pop failed to remove item")
	}

	check, exists := s.Peek()
	if !exists || check != 5 {
		t.Errorf("pop removed the wrong item")
	}
}

func TestPeek(t *testing.T) {
	s := New()
	want := 1
	s.Push(want)
	num, exists := s.Peek()
	if !exists {
		t.Error("item should exist but does no")
	}

	if num != want {
		t.Errorf("got %d want %d", num, want)
	}
}

func TestPeekEmpty(t *testing.T) {
	s := New()
	_, exists := s.Peek()
	if exists {
		t.Error("exists and should not")
	}
}

func TestLen(t *testing.T) {
	s := New()
	s.Push(1)
	want := 1
	l := s.Len()

	if l != want {
		t.Errorf("expected length to be %d got %d", want, l)
	}
}
