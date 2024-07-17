package generic

import (
	"log"
	"testing"
)

func TestSetInt(t *testing.T) {

	s := NewSet[int]()
	ix1 := s.Add(1)
	ix2 := s.Add(2)
	ix3 := s.Add(3)
	ix4 := s.Add(1)
	t.Log(ix1, ix2, ix3, ix4)

	s.Remove(0) // the first item
	count := s.Count()
	ix1 = s.Contains(1)
	ix2 = s.Contains(2)
	ix3 = s.Contains(3)
	ix4 = s.Contains(4)
	t.Logf("count = %d, ix1 = %d, ix2 = %d, ix3 = %d, ix4 = %d", count, ix1, ix2, ix3, ix4)

	if count != 2 || ix1 != -1 || ix2 != 0 || ix3 != 1 || ix4 != -1 {
		log.Fatal("invalid set operatios")
	}
}

type ts struct {
	i int
	v string
}

func TestSetStruct(t *testing.T) {

	ts1 := ts{1, "item"}
	ts2 := ts{2, "item"}
	ts3 := ts{3, "item"}
	ts4 := ts{1, "item"}
	t.Log(ts1, ts2, ts3, ts4)
	s := NewSet[ts]()
	ix1 := s.Add(ts1)
	ix2 := s.Add(ts2)
	ix3 := s.Add(ts3)
	ix4 := s.Add(ts{1, "item"})
	t.Log(ix1, ix2, ix3, ix4)

	s.Remove(0)
	count := s.Count()
	ix1 = s.Contains(ts1)
	ix2 = s.Contains(ts2)
	ix3 = s.Contains(ts3)
	ix4 = s.Contains(ts4)
	t.Logf("count = %d, ix1 = %d, ix2 = %d, ix3 = %d, ix4 = %d", count, ix1, ix2, ix3, ix4)

	if count != 2 || ix1 != -1 || ix2 != 0 || ix3 != 1 || ix4 != -1 {
		log.Fatal("invalid set operatios")
	}
}
