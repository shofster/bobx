package strings

import "testing"

func TestSoundex(t *testing.T) {

	var sl = StringList{
		"item1",
		"item2",
		"item3",
	}
	t.Log("Prepend item0, no max")
	rsl := Prepend(sl, "item0", 0)
	if rsl[0] != "item0" || len(rsl) != 4 {
		t.Errorf("expected {item0, item1, item2, item3}, got %v", rsl)
	}
	t.Log("Prepend item0, max 3")
	rsl = Prepend(sl, "item0", 3)
	if rsl[0] != "item0" || len(rsl) != 3 {
		t.Errorf("expected {item0, item1, item2}, got %v", rsl)
	}

	t.Log("Find item2")
	f := Find(sl, func(i int, s string) (b bool) {
		if i == 1 && s == "item2" {
			b = true
		}
		return
	})
	if !f {
		t.Errorf("expected true, got %v", f)
	}

	t.Log("Recent item3, max 3")
	rsl = Recent(sl, "item3", 3)
	if rsl[0] != "item3" || len(rsl) != 3 {
		t.Errorf("expected {item3, item1, item2}, got %v", rsl)
	}

	t.Log("Remove item2")
	rsl = Remove(sl, "item2")
	if rsl[0] != "item1" || rsl[1] != "item3" || len(rsl) != 2 {
		t.Errorf("expected {item1, item3}, got %v", rsl)
	}

}
