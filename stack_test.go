package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New[string]()
	s.Push("John")
	s.Push("Jane")

	result1, ok1 := s.Pop()
	if result1 != "Jane" && ok1 != true {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result1, "Jane")
	}

	result2, ok2 := s.Peek()
	if result2 != "John" && ok2 != true {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result2, "John")
	}

	result3 := s.Len()
	if result3 != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result3, 1)
	}

	s.Pop()
	result4 := s.IsEmpty()
	if !result4 {
		t.Errorf("Result was incorrect, got: %t, want: %t.", result4, true)
	}

}
