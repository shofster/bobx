package generic

import "testing"

type stackStruct struct {
	ix int
	v  string
}

func TestStructStack(t *testing.T) {
	s := NewStack[stackStruct]()
	s.Push(stackStruct{0, "a"})
	s.Push(stackStruct{1, "b"})
	v, b := s.Pop()
	if !b || v.ix != 1 {
		t.Errorf("pop; stackstruct 1 got = %#v", v)
	}
	v, b = s.Pop()
	if !b || v.ix != 0 {
		t.Errorf("pop; stackstruct 0 got = %#v", v)
	}
}

func TestStringStack(t *testing.T) {
	s := NewStack[string]()
	s.Push("item a")
	s.Push("item b")
	s.Push("item c")
	s.Clear()
	s.Push("item 3")
	s.Push("item 2")
	s.Push("item 1")
	item, b := s.Peek()
	if !b || item != "item 1" {
		t.Errorf("peek; item 1 got = %s", item)
	}
	item, b = s.Pop()
	if !b || item != "item 1" {
		t.Errorf("pop; item 1 got = %s", item)
	}
	item, b = s.Pop()
	if !b || item != "item 2" {
		t.Errorf("pop; item 2 got = %s", item)
	}
	item, b = s.Pop()
	if !b || item != "item 3" {
		t.Errorf("pop; item 3 got = %s", item)
	}
	n := s.Count()
	if n != 0 {
		t.Error("Stack wasn't empty")
	}
	item, b = s.Pop()
	if b || item != "" {
		t.Errorf("empty pop; got = %s", item)
	}
}

func TestStackReverse(t *testing.T) {
	s := NewStack[string]()
	s.Push("item a")
	s.Push("item b")
	c := s.Push("item c")
	if s.Count() != 3 || c != 3 {
		t.Errorf("Expected Stack size 3 got %d %d", s.Count(), c)
	}
	n := 0
	s.Reverse(func(s string) {
		switch n {
		case 0:
			if s != "item a" {
				t.Errorf("Iterate item %d failed: \"item a\" \"%s\"", n, s)
			}
		case 1:
			if s != "item b" {
				t.Errorf("Iterate item %d failed: \"item b\" \"%s\"", n, s)
			}
		case 2:
			if s != "item c" {
				t.Errorf("Iterate item %d failed: \"item c\" \"%s\"", n, s)
			}
		}
		n++
	})
}
