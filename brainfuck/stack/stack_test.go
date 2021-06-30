package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("wrong stack length, got %d", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	pop1, err := s.Pop()
    if err != nil {
        t.Fatalf("parsing error %+v", err)
    }
	if pop1 != 3 {
		t.Errorf("wrong value, got %d", pop1)
	}

	pop2, err := s.Pop()
    if err != nil {
        t.Fatalf("parsing error %+v", err)
    }
	if pop2 != 2 {
		t.Errorf("wrong value, got %d", pop2)
	}
}

func TestPushPopEmpty(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

    s.Pop()
    s.Pop()
    s.Pop()

    if s.Size() != 0 {
        t.Fatalf("invalid stack size. want=%d got=%d", 0, s.Size())
    }
}